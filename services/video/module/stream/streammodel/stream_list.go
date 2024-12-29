package streammodel

import (
	"github.com/caovanhoang63/hiholive/shared/golang/core"
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
	CurrentView    int       `json:"currentView" gorm:"-"`
	IsRerun        bool      `json:"isRerun" gorm:"column:is_rerun"`
	State          string    `json:"state" gorm:"column:state"`
}

func (s *StreamList) Mask() {
	s.BaseModel.Mask(core.DbTypeStream)
	s.ChannelFakeId = core.NewUIDP(uint32(s.ChannelId), core.DbTypeChannel, 0)
	s.CategoryFakeId = core.NewUIDP(uint32(s.CategoryId), core.DbTypeCategory, 0)
	if s.Category != nil {
		s.Category.Mask(core.DbTypeCategory)
	}
	if s.Channel != nil {
		s.Channel.Mask()
	}
}
