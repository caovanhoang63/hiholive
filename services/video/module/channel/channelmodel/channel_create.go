package channelmodel

import "github.com/caovanhoang63/hiholive/shared/go/core"

type ChannelCreate struct {
	UserId      int         `json:"-" gorm:"column:user_id"`
	UserUID     *core.UID   `json:"userId" gorm:"-"`
	Panel       *core.Image `json:"panel" gorm:"column:panel"`
	Description string      `json:"description" gorm:"column:description"`
	Url         string      `json:"url" gorm:"column:url"`
	Contact     string      `json:"contact" gorm:"column:contact"`
}

func (c *ChannelCreate) UnMask() {
	c.UserId = int(c.UserUID.GetLocalID())
}
