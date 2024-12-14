package stmodel

import "github.com/caovanhoang63/hiholive/shared/go/core"

type Setting struct {
	core.BaseModel `bson:",inline"`
	Name           string `json:"name" form:"name" gorm:"column:name"`
	Value          string `json:"value" form:"value" gorm:"column:value"`
	Status         int    `json:"status" form:"status" gorm:"column:status"`
}

func (Setting) TableName() string {
	return "system_settings"
}