package stmodel

type SettingCreate struct {
	Name  string `json:"name" gorm:"column:name" validate:"required"`
	Value string `json:"value" gorm:"column:value"`
}

type SettingUpdate struct {
	Name  string `json:"name" form:"name" gorm:"column:name"`
	Value any    `json:"value" form:"value" gorm:"column:value" validate:"required"`
}
