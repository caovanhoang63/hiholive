package channelmodel

import "github.com/caovanhoang63/hiholive/shared/golang/core"

type ChannelUpdate struct {
	Panel       *core.Image `json:"panel" gorm:"column:panel"`
	Image       *core.Image `json:"image" gorm:"column:image"`
	Description string      `json:"description" gorm:"column:description"`
	Contact     string      `json:"contact" gorm:"column:contact"`
}
