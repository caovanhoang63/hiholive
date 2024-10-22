package entity

import "hiholive/shared/go/core"

type UserCreate struct {
	core.BaseModel `json:",inline"`
	Email          string `json:"email" gorm:"column:email"`
	FirstName      string `json:"first_name" gorm:"column:first_name"`
	LastName       string `json:"last_name" gorm:"column:last_name"`
	DisplayName    string `json:"display_name" gorm:"column:display_name"`
	Gender         Gender `json:"gender" gorm:"column:gender"`
	Password       string `json:"password" gorm:"column:password"`
	Salt           string `json:"salt" gorm:"column:salt"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserUpdate struct {
	PhoneNumber string      `json:"phone_number" gorm:"column:phone_number"`
	Address     string      `json:"address" gorm:"column:address"`
	FirstName   string      `json:"first_name" gorm:"column:first_name"`
	LastName    string      `json:"last_name" gorm:"column:last_name"`
	DisplayName string      `json:"display_name" gorm:"column:display_name"`
	DateOfBirth string      `json:"date_of_birth" gorm:"column:date_of_birth"`
	Gender      Gender      `json:"gender" gorm:"column:gender"`
	Avatar      *core.Image `json:"avatar" gorm:"column:avatar"`
	Bio         string      `json:"bio" gorm:"column:bio"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}
