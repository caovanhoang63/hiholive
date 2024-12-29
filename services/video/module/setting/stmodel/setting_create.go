package stmodel

import "github.com/caovanhoang63/hiholive/shared/golang/core"

type SettingCreate struct {
	core.BaseModel `bson:",inline"`
	Name           string `json:"name" gorm:"column:name" validate:"required"`
	Value          any    `gorm:"serializer:json" json:"value" form:"value" validate:"required"`
}

type SettingUpdate struct {
	Name  *string `json:"name" form:"name" gorm:"column:name"`
	Value any     `gorm:"serializer:json" json:"value" form:"value" validate:"required"`
}
