package rtmpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/caovanhoang63/hiholive/shared/go/asyncjob"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/pubsub"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	flvtag "github.com/yutopp/go-flv/tag"
	"github.com/yutopp/go-rtmp"
	rtmpmsg "github.com/yutopp/go-rtmp/message"
	"golang.org/x/net/context"
	"io"
	"time"
)

//	HLS 720p
//
// Client -> OBS -> RTMP -> FFMPEG -> HLS 1080p    -> Cloudfront -> videojs
//
//	HLS 440p

type HlsClient interface {
	NewHlsStream(ctx context.Context, streamId, serverUrl, streamKey string, fps, resolution int) (err error)
}

var curStream = map[string]core.StreamState{}

var _ rtmp.Handler = (*Handler)(nil)

// Handler An RTMP connection handler
type Handler struct {
	rtmp.DefaultHandler
	relayService *RelayService
	logger       srvctx.Logger
	// conn represents the RTMP connection associated with this handler.
	conn      *rtmp.Conn
	rdClient  *redis.Client
	hlsClient HlsClient
	Stream    *core.StreamState
	// pub represents the publishing entity, handling media streams such as audio, video, and metadata for RTMP connections.
	pub             *Pub
	streamState     string
	cancelErrorFunc context.CancelFunc
	// sub represents a subscriber for handling events or media streams during RTMP playback sessions.
	sub *Sub

	ps pubsub.Pubsub
}

// NewHandler creates and returns a new Handler instance for managing RTMP connections with provided dependencies.
func NewHandler(relayService *RelayService, rd *redis.Client, hlsClient HlsClient, ps pubsub.Pubsub) *Handler {
	return &Handler{
		logger:       srvctx.DefaultLogger,
		relayService: relayService,
		rdClient:     rd,
		hlsClient:    hlsClient,
		ps:           ps,
	}

}

// OnServe initializes the RTMP connection for the handler and assigns it to the handler's 'conn' field.
func (h *Handler) OnServe(conn *rtmp.Conn) {
	h.conn = conn
}

// OnConnect validates the application name during an RTMP connection request and initializes the connection process.
func (h *Handler) OnConnect(timestamp uint32, cmd *rtmpmsg.NetConnectionConnect) error {
	h.logger.WithField(srvctx.Field{"cmd": cmd}).Info("OnConnect")
	if cmd.Command.App != core.StreamDomain {
		return errors.New("OnConnect: Invalid App Name")
	}
	return nil
}

func (h *Handler) OnCreateStream(timestamp uint32, cmd *rtmpmsg.NetConnectionCreateStream) error {
	h.logger.WithField(srvctx.Field{"cmd": cmd}).Info("OnCreateStream")
	return nil
}

// OnPublish handles publishing requests for incoming RTMP streams, validates publishing parameters, and initializes resources.
func (h *Handler) OnPublish(ctx *rtmp.StreamContext, timestamp uint32, cmd *rtmpmsg.NetStreamPublish) error {
	h.logger.WithField(srvctx.Field{"cmd": cmd}).Info("OnPublish")
	if h.sub != nil {
		return errors.New("Cannot publish to this stream")
	}

	if cmd.PublishingName == "" {
		return errors.New("PublishingName is empty")
	}
	pubsub, err := h.relayService.NewPubsub(cmd.PublishingName)
	if err != nil {
		return errors.Wrap(err, "Failed to create pubsub")
	}
	byteData, err := h.rdClient.Get(context.Background(), fmt.Sprintf("streamKey:%s", cmd.PublishingName)).Result()

	var streamInfo core.StreamState
	_ = json.Unmarshal([]byte(byteData), &streamInfo)

	streamInfo.StreamKey = cmd.PublishingName
	if errors.Is(err, redis.Nil) || err != nil {
		log.Print(err)
		return errors.New("PublishingName does not exist")
	}

	h.Stream = &streamInfo

	pub := pubsub.Pub()

	h.pub = pub

	return nil
}

// OnPlay sets up a subscriber for RTMP playback on a specified stream and validates the stream's availability.
func (h *Handler) OnPlay(ctx *rtmp.StreamContext, timestamp uint32, cmd *rtmpmsg.NetStreamPlay) error {
	if h.sub != nil {
		return errors.New("Cannot play on this stream")
	}

	pubsub, err := h.relayService.GetPubsub(cmd.StreamName)
	if err != nil {
		return errors.Wrap(err, "Failed to get pubsub")
	}

	sub := pubsub.Sub()
	sub.eventCallback = onEventCallback(h.conn, ctx.StreamID)

	h.sub = sub

	return nil
}

// OnSetDataFrame processes the "onMetaData" script data from a NetStreamSetDataFrame message and publishes it to subscribers.
func (h *Handler) OnSetDataFrame(timestamp uint32, data *rtmpmsg.NetStreamSetDataFrame) error {
	r := bytes.NewReader(data.Payload)
	var script flvtag.ScriptData
	if err := flvtag.DecodeScriptData(r, &script); err != nil {
		h.logger.Errorf("Failed to decode script data: Err = %+v", err)
		return nil // ignore
	}
	h.logger.WithField(srvctx.Field{"Script": script}).Info("SetDataFrame")

	object := script.Objects["onMetaData"]
	fps := object["framerate"].(float64)
	height := object["height"].(float64)

	address, err := core.GetServerAddress()
	if err != nil {
		h.logger.Error(err)
		return err
	}

	// Format URL động
	serverUrl := fmt.Sprintf("rtmp://%s:1935/stream", address)

	id, err := core.FromBase58(h.Stream.Uid)

	if err != nil {
		return core.ErrInternalServerError.WithWrap(err)
	}

	go func() {
		core.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return h.hlsClient.NewHlsStream(ctx, h.Stream.Uid, serverUrl, h.Stream.StreamKey, int(fps), int(height))
		})

		// Retry 3 time to call to hls server
		job.SetRetryDurations([]time.Duration{
			time.Second * 1,
			time.Second * 2,
			time.Second * 4,
			time.Second * 8,
		})
		if err = job.RunWithRetry(context.Background()); err != nil {
			h.logger.Error(err)
		}
	}()
	if err = h.ps.Publish(context.Background(), core.TopicStreamStart, pubsub.NewMessage(id)); err != nil {
		h.logger.Error(err)
	}
	if h.streamState == "error" {
		h.cancelErrorFunc()
	}
	h.streamState = "started"
	_ = h.pub.Publish(&flvtag.FlvTag{
		TagType:   flvtag.TagTypeScriptData,
		Timestamp: timestamp,
		Data:      &script,
	})

	return nil
}

// OnAudio processes incoming audio data, decodes it, and publishes it as an FLV tag for subscribers.
func (h *Handler) OnAudio(timestamp uint32, payload io.Reader) error {
	var audio flvtag.AudioData
	if err := flvtag.DecodeAudioData(payload, &audio); err != nil {
		h.logger.Errorf("Failed to decode audio data: Err = %+v", err)
		return err
	}

	flvBody := new(bytes.Buffer)
	if _, err := io.Copy(flvBody, audio.Data); err != nil {
		h.logger.Errorf("Failed to copy audio data: Err = %+v", err)
		return err
	}
	audio.Data = flvBody

	_ = h.pub.Publish(&flvtag.FlvTag{
		TagType:   flvtag.TagTypeAudio,
		Timestamp: timestamp,
		Data:      &audio,
	})

	return nil
}

// OnVideo processes incoming video data, decodes it, and publishes it as an FLV tag for subscribers.
func (h *Handler) OnVideo(timestamp uint32, payload io.Reader) error {
	var video flvtag.VideoData
	if err := flvtag.DecodeVideoData(payload, &video); err != nil {
		h.logger.Errorf("Failed to decode video data: Err = %+v", err)
		return err
	}

	// Need deep copy because payload will be recycled
	flvBody := new(bytes.Buffer)
	if _, err := io.Copy(flvBody, video.Data); err != nil {
		h.logger.Errorf("Failed to copy video data: Err = %+v", err)
		return err
	}
	video.Data = flvBody

	_ = h.pub.Publish(&flvtag.FlvTag{
		TagType:   flvtag.TagTypeVideo,
		Timestamp: timestamp,
		Data:      &video,
	})

	return nil
}

// OnClose cleans up resources associated with the handler by closing the publisher and subscriber, if they are initialized.
func (h *Handler) OnClose() {
	h.logger.Infof("OnClose")
	if h.pub != nil {
		h.handleEndStream()
		_ = h.pub.Close()
	}

	if h.sub != nil {
		fmt.Println("Sub close ")
		_ = h.sub.Close()
	}
}

func (h *Handler) OnError(e error) {
	fmt.Println("OnError:", e)
	ctx, cancel := context.WithCancel(context.Background())
	h.cancelErrorFunc = cancel
	go func() {
		core.AppRecover()
		h.streamState = "error"

		select {
		// Wait 3 minute after stop stream
		case <-time.After(time.Minute * 3):
			h.handleEndStream()
			fmt.Println("OnError")
		case <-ctx.Done(): // Nếu context bị hủy
			fmt.Println("Error handling was canceled.")
		}
	}()
	fmt.Println(e)
}

func (h *Handler) OnStop() {

}

func (h *Handler) handleEndStream() {
	go func() {
		core.AppRecover()
		id, _ := core.FromBase58(h.Stream.Uid)
		_ = h.ps.Publish(context.Background(), core.TopicStreamEnded, pubsub.NewMessage(map[string]interface{}{
			"stream_id": id,
			"timestamp": time.Now(),
		}))
		h.rdClient.Del(context.Background(), fmt.Sprintf("streamKey:%s", h.Stream.StreamKey)).Result()
	}()

}
