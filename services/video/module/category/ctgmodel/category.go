package ctgmodel

import "github.com/caovanhoang63/hiholive/shared/go/core"

type Category struct {
	core.BaseModel `json:",inline"`
	Name           string      `json:"name,omitempty" gorm:"column:name"`
	Description    string      `json:"description,omitempty" gorm:"column:description"`
	Image          *core.Image `json:"image,omitempty" gorm:"column:image"`
	Status         int         `json:"status,omitempty" gorm:"column:status"`
}

func (c Category) TableName() string {
	return "categories"
}

func (c *Category) Mask() {
	c.BaseModel.Mask(core.DbTypeCategory)
}
