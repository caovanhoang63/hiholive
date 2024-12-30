package channelmodel

import "github.com/caovanhoang63/hiholive/shared/golang/core"

type ChannelCreate struct {
	core.BaseModel `bson:",inline"`
	UserId         int         `json:"-" gorm:"column:user_id"`
	UserUID        *core.UID   `json:"userId" gorm:"-"`
	DisplayName    string      `json:"displayName" gorm:"column:display_name" validate:"required"`
	UserName       string      `json:"userName" gorm:"column:user_name" validate:"required"`
	Image          *core.Image `json:"image" gorm:"column:image"`
	Panel          *core.Image `json:"panel" gorm:"column:panel"`
	Description    string      `json:"description" gorm:"column:description"`
	Contact        string      `json:"contact" gorm:"column:contact"`
}

func (c *ChannelCreate) UnMask() {
	c.UserId = int(c.UserUID.GetLocalID())
}
