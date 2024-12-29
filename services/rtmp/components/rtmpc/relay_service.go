package rtmpc

import (
	"errors"
	"fmt"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"golang.org/x/net/context"
	"sync"
)

// RelayService TODO: Create this service per apps.
// In this example, this instance is singleton.
type RelayService struct {
	streams     map[string]*Pubsub
	streamState map[string]*core.StreamState
	cancel      map[string]*context.CancelFunc
	m           sync.Mutex
}

func NewRelayService() *RelayService {
	return &RelayService{
		streams:     make(map[string]*Pubsub),
		streamState: make(map[string]*core.StreamState),
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
