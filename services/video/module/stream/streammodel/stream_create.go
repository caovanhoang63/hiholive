package streammodel

import (
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/google/uuid"
	"time"
)

type StreamCreate struct {
	core.BaseModel     `json:",inline"`
	Title              string     `json:"title" gorm:"column:title" validate:"required"`
	Description        string     `json:"description" gorm:"column:description" validate:"required"`
	ChannelId          int        `json:"-" gorm:"column:channel_id"`
	ChannelFakeId      *core.UID  `json:"ChannelId" gorm:"-"`
	CategoryId         int        `json:"-" gorm:"column:category_id"`
	CategoryFakeId     *core.UID  `json:"categoryId" gorm:"-"`
	IsRerun            bool       `json:"isRerun" gorm:"column:is_rerun"`
	Notification       string     `json:"notification" gorm:"column:notification"`
	ScheduledStartTime *time.Time `json:"scheduledStartTime" gorm:"column:scheduled_start_time"`
	StreamKey          *uuid.UUID `json:"streamKey" gorm:"column:stream_key"`
}

func (s *StreamCreate) UnMask() {
	s.CategoryId = int(s.CategoryFakeId.GetLocalID())
	s.ChannelId = int(s.ChannelFakeId.GetLocalID())
}

type StreamCreateResponse struct {
	StreamKey *uuid.UUID `json:"streamKey" gorm:"column:stream_key"`
	RtmpLink  string     `json:"rtmpLink" gorm:"-"`
}
