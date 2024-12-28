package channelmodel

import "github.com/caovanhoang63/hiholive/shared/go/core"

type Channel struct {
	core.BaseModel `bson:",inline"`
	UserId         int         `json:"-" gorm:"column:user_id"`
	UserUID        *core.UID   `json:"userId" gorm:"-"`
	Panel          *core.Image `json:"panel" gorm:"column:panel"`
	Image          *core.Image `json:"image" gorm:"column:image"`
	UserName       string      `json:"userName" gorm:"column:user_name"`
	DisplayName    string      `json:"displayName" gorm:"column:display_name"`
	Description    string      `json:"description" gorm:"column:description"`
	Url            string      `json:"url" gorm:"column:url"`
	Contact        string      `json:"contact" gorm:"column:contact"`
	Status         int         `json:"status" gorm:"column:status"`
}

func (Channel) TableName() string {
	return "channels"
}

func (channel *Channel) Mask() {
	channel.BaseModel.Mask(core.DbTypeChannel)
	channel.UserUID = core.NewUIDP(uint32(channel.Id), core.DbTypeUser, 0)
}
