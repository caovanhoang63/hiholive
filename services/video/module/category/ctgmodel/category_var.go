package ctgmodel

import "github.com/caovanhoang63/hiholive/shared/go/core"

type CategoryCreate struct {
	core.BaseModel `json:",inline"`
	Name           string      `json:"name,omitempty" gorm:"column:name" validate:"required"`
	Description    string      `json:"description,omitempty" gorm:"column:description"`
	Image          *core.Image `json:"image,omitempty" gorm:"column:image"`
}

type CategoryUpdate struct {
	Name        string      `json:"name,omitempty" gorm:"column:name"`
	Description string      `json:"description,omitempty" gorm:"column:description"`
	Image       *core.Image `json:"image,omitempty" gorm:"column:image"`
}
