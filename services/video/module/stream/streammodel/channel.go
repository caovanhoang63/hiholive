package streammodel

import "github.com/caovanhoang63/hiholive/shared/golang/core"

type Channel struct {
	core.BaseModel `bson:",inline"`
	UserId         int         `json:"-" gorm:"column:user_id"`
	UserUID        *core.UID   `json:"userId" gorm:"-"`
	Image          *core.Image `json:"image" gorm:"column:image"`
	UserName       string      `json:"userName" gorm:"column:user_name"`
	DisplayName    string      `json:"displayName" gorm:"column:display_name"`
}

func (channel *Channel) Mask() {
	channel.BaseModel.Mask(core.DbTypeChannel)
	channel.UserUID = core.NewUIDP(uint32(channel.UserId), core.DbTypeUser, 0)
}
