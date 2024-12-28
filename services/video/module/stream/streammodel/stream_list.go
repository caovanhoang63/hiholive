package streammodel

import (
	"github.com/caovanhoang63/hiholive/shared/go/core"
)

type StreamList struct {
	core.BaseModel `json:",inline"`
	ChannelId      int       `json:"-" gorm:"column:channel_id"`
	ChannelFakeId  *core.UID `json:"ChannelId" gorm:"-"`
	Channel        *Channel  `json:"channel" gorm:"foreignkey:ChannelId"`
	Title          string    `json:"title" gorm:"column:title"`
	Description    string    `json:"description" gorm:"column:description"`
	CategoryId     int       `json:"-" gorm:"column:category_id"`
	Category       *Category `json:"category" gorm:"foreignkey:CategoryId;preload=false"`
	CategoryFakeId *core.UID `json:"-" gorm:"-"`
	IsRerun        bool      `json:"isRerun" gorm:"column:is_rerun"`
	State          string    `json:"state" gorm:"column:state"`
}
