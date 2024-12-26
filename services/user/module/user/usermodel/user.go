package usermodel

import (
	"github.com/caovanhoang63/hiholive/shared/go/core"
)

type SystemRole string
type Gender string

const (
	RoleAdmin     string = "admin"
	RoleStreamer  string = "streamer"
	RoleUser      string = "viewer"
	RoleModerator string = "moderator"
)

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)

type User struct {
	core.BaseModel `json:",inline"`
	PhoneNumber    string      `json:"phone_number" gorm:"column:phone_number"`
	Address        string      `json:"address" gorm:"column:address"`
	FirstName      string      `json:"first_name" gorm:"column:first_name"`
	LastName       string      `json:"last_name" gorm:"column:last_name"`
	DisplayName    string      `json:"display_name" gorm:"column:display_name"`
	DateOfBirth    string      `json:"date_of_birth" gorm:"column:date_of_birth"`
	Gender         Gender      `json:"gender" gorm:"column:gender"`
	SystemRole     string      `json:"SystemRole" gorm:"column:system_role"`
	Avatar         *core.Image `json:"avatar" gorm:"column:avatar"`
	Bio            string      `json:"bio" gorm:"column:bio"`
	Status         int         `json:"status" gorm:"column:status"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask() {
	u.BaseModel.Mask(core.DbTypeUser)
}
