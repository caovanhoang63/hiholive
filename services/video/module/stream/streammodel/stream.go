package streammodel

import (
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/google/uuid"
	"time"
)

type Stream struct {
	core.BaseModel     `json:",inline"`
	ChannelId          int        `json:"-" gorm:"column:channel_id"`
	ChannelFakeId      *core.UID  `json:"ChannelId" gorm:"-"`
	Title              string     `json:"title" gorm:"column:title"`
	Notification       string     `json:"notification" gorm:"column:notification"`
	Description        string     `json:"description" gorm:"column:description"`
	CategoryId         int        `json:"-" gorm:"column:category_id"`
	CategoryFakeId     *core.UID  `json:"categoryId" gorm:"-"`
	IsRerun            bool       `json:"isRerun" gorm:"column:is_rerun"`
	StreamKey          *uuid.UUID `json:"streamKey" gorm:"column:stream_key"`
	State              string     `json:"state" gorm:"column:state"`
	ActualStartTime    *time.Time `json:"actualStartTime" gorm:"column:actual_start_time"`
	ActualEndTime      *time.Time `json:"actualEndTime" gorm:"column:actual_end_time"`
	PeakConcurrentView int        `json:"peakConcurrentView" gorm:"column:peak_concurrent_view"`
	TotalUniqueViewers int        `json:"totalUniqueViewers" gorm:"column:total_unique_viewers"`
	ScheduledStart     *time.Time `json:"scheduledStart" gorm:"column:scheduled_start"`
	Status             int        `json:"status" gorm:"column:status"`
}

func (Stream) TableName() string {
	return "live_streams"
}

func (s *Stream) Mask() {
	s.BaseModel.Mask(core.DbTypeStream)
	s.ChannelFakeId = core.NewUIDP(uint32(s.ChannelId), core.DbTypeChannel, 0)
	s.CategoryFakeId = core.NewUIDP(uint32(s.CategoryId), core.DbTypeCategory, 0)
}

const (
	StreamStateScheduled string = "scheduled"
	StreamStateRunning   string = "running"
	StreamStateEnded     string = "ended"
	StreamStatePending   string = "pending"
)
