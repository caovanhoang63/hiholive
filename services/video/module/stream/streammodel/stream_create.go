package streammodel

import (
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"time"
)

type StreamCreate struct {
	Title          string     `json:"title" gorm:"column:title" validate:"required"`
	Description    string     `json:"description" gorm:"column:description" validate:"required"`
	ChannelId      int        `json:"-" gorm:"column:channel_id"`
	ChannelFakeId  *core.UID  `json:"ChannelId" gorm:"-"`
	CategoryId     int        `json:"-" gorm:"column:category_id"`
	CategoryFakeId *core.UID  `json:"categoryId" gorm:"-"`
	IsRerun        bool       `json:"isRerun" gorm:"column:is_rerun"`
	State          string     `json:"state" gorm:"column:state"`
	Notification   string     `json:"notification" gorm:"column:notification"`
	ScheduledStart *time.Time `json:"scheduledStart" gorm:"column:scheduled_start"`
}

func (s *StreamCreate) UnMask() {
	s.CategoryId = int(s.CategoryFakeId.GetLocalID())
	s.ChannelId = int(s.ChannelFakeId.GetLocalID())
}
