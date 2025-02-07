package usermodel

import "github.com/caovanhoang63/hiholive/shared/golang/core"

type UserCreate struct {
	core.BaseModel `json:",inline"`
	Email          string `json:"email" gorm:"column:email"`
	FirstName      string `json:"first_name" gorm:"column:first_name"`
	LastName       string `json:"last_name" gorm:"column:last_name"`
	DisplayName    string `json:"display_name" gorm:"column:display_name"`
	UserName       string `json:"user_name" gorm:"column:user_name"`
	Gender         string `json:"gender" gorm:"column:gender"`
	SystemRole     string `json:"system_role" gorm:"column:system_role"`
}

func NewUserForCreation(firstName, lastName, email string) UserCreate {
	return UserCreate{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserUpdate struct {
	PhoneNumber string      `json:"phone_number" gorm:"column:phone_number"`
	Address     string      `json:"address" gorm:"column:address"`
	FirstName   string      `json:"first_name" gorm:"column:first_name"`
	LastName    string      `json:"last_name" gorm:"column:last_name"`
	DateOfBirth string      `json:"date_of_birth" gorm:"column:date_of_birth"`
	Gender      string      `json:"gender" gorm:"column:gender"`
	Avatar      *core.Image `json:"avatar" gorm:"column:avatar"`
}

type UserNameAndDisplayName struct {
	UserName    string `json:"user_name" gorm:"column:user_name" validate:"required"`
	DisplayName string `json:"display_name" gorm:"column:display_name" validate:"required"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}
