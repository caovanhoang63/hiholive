package entity

import "github.com/caovanhoang63/hiholive/shared/go/core"

type Auth struct {
	core.BaseModel
	UserId     int    `json:"userId" gorm:"column:user_id;" db:"user_id"`
	AuthType   string `json:"authType" gorm:"column:auth_type;" db:"auth_type"`
	Email      string `json:"email" gorm:"column:email;" db:"email"`
	Salt       string `json:"salt" gorm:"column:salt;" db:"salt"`
	Password   string `json:"password" gorm:"column:password;" db:"password"`
	FacebookId string `json:"facebookId" gorm:"column:facebook_id" db:"facebook_id"`
}

func (Auth) TableName() string { return "auths" }

func NewAuthWithEmailPassword(userId int, email, salt, password string) Auth {
	return Auth{
		BaseModel: core.NewBaseModel(),
		UserId:    userId,
		Email:     email,
		Salt:      salt,
		Password:  password,
		AuthType:  "email_password",
	}
}
