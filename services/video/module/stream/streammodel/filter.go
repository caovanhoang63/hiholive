package streammodel

import (
	"errors"
	"github.com/caovanhoang63/hiholive/shared/go/core"
)

type StreamFilter struct {
	State          string `json:"state" form:"state"`
	CategoryFakeId string `json:"categoryId" form:"categoryId"`
	CategoryId     int    `json:"-" form:"-"`
	ChannelFakeId  string `json:"channelId" form:"channelId"`
	ChannelId      int    `json:"-" form:"-"`
	UserName       string `json:"userName" form:"userName"`
	Title          string `json:"title" form:"title"`
}

func (s *StreamFilter) Process() error {
	states := []string{StreamStateRunning, StreamStateEnded, StreamStatePending, StreamStateScheduled}
	ok := false
	if s.State != "" {
		for _, state := range states {
			if state == s.State {
				ok = true
				continue
			}
		}
		if !ok {
			return errors.New("invalid stream state")
		}
	}

	if s.CategoryFakeId != "" {
		cUid, err := core.FromBase58(s.CategoryFakeId)
		if err != nil {
			return errors.New("invalid channel id")
		}
		s.CategoryId = int(cUid.GetLocalID())
	}

	if s.ChannelFakeId != "" {
		cUid, err := core.FromBase58(s.ChannelFakeId)
		if err != nil {
			return errors.New("invalid channel id")
		}
		s.ChannelId = int(cUid.GetLocalID())
	}

	return nil
}
