package streammodel

import (
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/google/uuid"
)

type StreamList struct {
	core.BaseModel `json:",inline"`
	ChannelId      int        `json:"-" gorm:"column:channel_id"`
	ChannelFakeId  *core.UID  `json:"ChannelId" gorm:"-"`
	Title          string     `json:"title" gorm:"column:title"`
	Notification   string     `json:"notification" gorm:"column:notification"`
	Description    string     `json:"description" gorm:"column:description"`
	CategoryId     int        `json:"-" gorm:"column:category_id"`
	Category       *Category  `json:"category" gorm:"foreignkey:CategoryId;preload=false"`
	CategoryFakeId *core.UID  `json:"-" gorm:"-"`
	IsRerun        bool       `json:"isRerun" gorm:"column:is_rerun"`
	StreamKey      *uuid.UUID `json:"streamKey" gorm:"column:stream_key"`
	State          string     `json:"state" gorm:"column:state"`
	Status         int        `json:"status" gorm:"column:status"`
}
