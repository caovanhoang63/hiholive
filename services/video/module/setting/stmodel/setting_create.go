package stmodel

type SettingCreate struct {
	Name  string                 `json:"name" gorm:"column:name" validate:"required"`
	Value map[string]interface{} `gorm:"serializer:json"  validate:"required"`
}

type SettingUpdate struct {
	Name  string                 `json:"name" form:"name" gorm:"column:name"`
	Value map[string]interface{} `gorm:"serializer:json" validate:"required"`
}
