package rtmpc

import (
	"errors"
	"fmt"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"sync"

	"time"
)

// RelayService TODO: Create this service per apps.
// In this example, this instance is singleton.
type RelayService struct {
	streams     map[string]*Pubsub
	streamState map[string]*core.StreamState
	cancel      map[string]*context.CancelFunc
	m           sync.Mutex
	ps          pubsub.Pubsub
	rd          *redis.Client
}

func NewRelayService(ps pubsub.Pubsub) *RelayService {
	return &RelayService{
		streams:     make(map[string]*Pubsub),
		streamState: make(map[string]*core.StreamState),
		ps:          ps,
	}
}

var ErrAlreadyPublished = errors.New("already published")

func (s *RelayService) CancelError(streamKey string) bool {
	s.m.Lock()
	defer s.m.Unlock()
	cancel, ok := s.cancel[streamKey]
	if ok {
		(*cancel)()
	}
	return ok
}

func (s *RelayService) OnError(streamKey string, e error) {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel[streamKey] = &cancel
	stream := s.streamState[streamKey]
	go func(ctx context.Context) {
		defer core.AppRecover()
		stream.State = "error"
		select {
		// Wait 3 minute after stop stream
		case <-time.After(time.Minute * 3):
			go func() {
				defer core.AppRecover()
				id, _ := core.FromBase58(stream.Uid)
				_ = s.ps.Publish(context.Background(), core.TopicStreamEnded, pubsub.NewMessage(map[string]interface{}{
					"stream_id": id,
					"timestamp": time.Now(),
				}))
				s.rd.Del(context.Background(), fmt.Sprintf("streamKey:%s", streamKey)).Result()
			}()
			fmt.Println("OnError")
		case <-ctx.Done(): // Context cancelled ( Streamer reconnect to server)
			fmt.Println("Error handling was canceled.")
		}
	}(ctx)
}

func (s *RelayService) GetStream(streamKey string) (*core.StreamState, error) {
	s.m.Lock()
	defer s.m.Unlock()
	stream, ok := s.streamState[streamKey]
	if !ok {
		return nil, errors.New("stream not found")
	}
	return stream, nil
}

func (s *RelayService) NewPubsub(key string) (*Pubsub, error) {
	s.m.Lock()
	defer s.m.Unlock()

	if pubsub, ok := s.streams[key]; ok {
		return pubsub, ErrAlreadyPublished
	}

	pubsub := NewPubsub(s, key)

	s.streams[key] = pubsub

	return pubsub, nil
}

func (s *RelayService) GetPubsub(key string) (*Pubsub, error) {
	s.m.Lock()
	defer s.m.Unlock()

	pubsub, ok := s.streams[key]
	if !ok {
		return nil, fmt.Errorf("Not published: %s", key)
	}

	return pubsub, nil
}

func (s *RelayService) RemovePubsub(key string) error {
	s.m.Lock()
	defer s.m.Unlock()

	if _, ok := s.streams[key]; !ok {
		return fmt.Errorf("Not published: %s", key)
	}

	delete(s.streams, key)

	return nil
}
