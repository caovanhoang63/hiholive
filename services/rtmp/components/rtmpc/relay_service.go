package rtmpc

import (
	"errors"
	"fmt"
	"sync"
)

// RelayService TODO: Create this service per apps.
// In this example, this instance is singleton.
type RelayService struct {
	streams map[string]*Pubsub
	m       sync.Mutex
}

func NewRelayService() *RelayService {
	return &RelayService{
		streams: make(map[string]*Pubsub),
	}
}

var ErrAlreadyPublished = errors.New("already published")

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
