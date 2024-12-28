package streammodel

import "github.com/caovanhoang63/hiholive/shared/go/core"

type Channel struct {
	core.BaseModel `bson:",inline"`
	UserId         int         `json:"-" gorm:"column:user_id"`
	UserUID        *core.UID   `json:"userId" gorm:"-"`
	Panel          *core.Image `json:"panel" gorm:"column:panel"`
	Description    string      `json:"description" gorm:"column:description"`
	Url            string      `json:"url" gorm:"column:url"`
	Contact        string      `json:"contact" gorm:"column:contact"`
	Status         int         `json:"status" gorm:"column:status"`
}
