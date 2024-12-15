package stmodel

import (
	"github.com/caovanhoang63/hiholive/shared/go/core"
)

type Setting struct {
	core.BaseModel `bson:",inline"`
	Name           string `json:"name" form:"name" gorm:"column:name"`
	Value          any    `gorm:"serializer:json" json:"value" form:"value"`
	Status         int    `json:"status" form:"status" gorm:"column:status"`
}

func (Setting) TableName() string {
	return "system_settings"
}
