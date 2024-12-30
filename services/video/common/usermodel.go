package common

import "github.com/caovanhoang63/hiholive/shared/golang/core"

type User struct {
	Uid         *core.UID   `json:"id"`
	FirstName   string      `json:"first_name" gorm:"column:first_name"`
	LastName    string      `json:"last_name" gorm:"column:last_name"`
	UserName    string      `json:"user_name" gorm:"column:user_name"`
	DisplayName string      `json:"display_name" gorm:"column:display_name"`
	SystemRole  string      `json:"SystemRole" gorm:"column:system_role"`
	Avatar      *core.Image `json:"avatar" gorm:"column:avatar"`
}
