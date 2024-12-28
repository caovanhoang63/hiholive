package ctgmodel

import "github.com/caovanhoang63/hiholive/shared/go/core"

type Category struct {
	core.BaseModel `json:",inline"`
	Name           string      `json:"name" gorm:"column:name"`
	Description    string      `json:"description" gorm:"column:description"`
	Image          *core.Image `json:"image" gorm:"column:image"`
	Status         int         `json:"status" gorm:"column:status"`
	TotalContent   int         `json:"totalContent" gorm:"column:total_content"`
}

func (c Category) TableName() string {
	return "categories"
}

func (c *Category) Mask() {
	c.BaseModel.Mask(core.DbTypeCategory)
}
