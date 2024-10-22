package rtmpc

import (
	"bytes"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/pkg/errors"
	flvtag "github.com/yutopp/go-flv/tag"
	"github.com/yutopp/go-rtmp"
	rtmpmsg "github.com/yutopp/go-rtmp/message"

	"io"
)

//
//                                    HLS 720p
// Client -> OBS -> RTMP -> FFMPEG -> HLS 1080p    -> Cloudfront -> videojs
//                                    HLS 440p
//

var _ rtmp.Handler = (*Handler)(nil)

// Handler An RTMP connection handler
type Handler struct {
	rtmp.DefaultHandler
	relayService *RelayService
	logger       srvctx.Logger
	//
	conn *rtmp.Conn
	//
	pub *Pub
	sub *Sub
}

func NewHandler(relayService *RelayService) *Handler {
	return &Handler{
		logger:       srvctx.DefaultLogger,
		relayService: relayService,
	}
}

func (h *Handler) OnServe(conn *rtmp.Conn) {
	h.conn = conn
}

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

func (h *Handler) OnPublish(_ *rtmp.StreamContext, timestamp uint32, cmd *rtmpmsg.NetStreamPublish) error {
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

	if cmd.PublishingName != "test" {
		return errors.New("PublishingName is empty")
	}
	h.logger.Infof("KEY STREAM %s", cmd.PublishingName)

	pub := pubsub.Pub()

	h.pub = pub

	return nil
}

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

func (h *Handler) OnSetDataFrame(timestamp uint32, data *rtmpmsg.NetStreamSetDataFrame) error {
	r := bytes.NewReader(data.Payload)

	var script flvtag.ScriptData
	if err := flvtag.DecodeScriptData(r, &script); err != nil {
		h.logger.Errorf("Failed to decode script data: Err = %+v", err)
		return nil // ignore
	}

	h.logger.WithField(srvctx.Field{"Script": script}).Info("SetDataFrame")

	_ = h.pub.Publish(&flvtag.FlvTag{
		TagType:   flvtag.TagTypeScriptData,
		Timestamp: timestamp,
		Data:      &script,
	})

	return nil
}

func (h *Handler) OnAudio(timestamp uint32, payload io.Reader) error {
	var audio flvtag.AudioData
	if err := flvtag.DecodeAudioData(payload, &audio); err != nil {
		return err
	}

	flvBody := new(bytes.Buffer)
	if _, err := io.Copy(flvBody, audio.Data); err != nil {
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

func (h *Handler) OnVideo(timestamp uint32, payload io.Reader) error {
	var video flvtag.VideoData
	if err := flvtag.DecodeVideoData(payload, &video); err != nil {
		return err
	}

	// Need deep copy because payload will be recycled
	flvBody := new(bytes.Buffer)
	if _, err := io.Copy(flvBody, video.Data); err != nil {
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

func (h *Handler) OnClose() {
	h.logger.Infof("OnClose")

	if h.pub != nil {
		_ = h.pub.Close()
	}

	if h.sub != nil {
		_ = h.sub.Close()
	}
}
