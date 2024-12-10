package core

import "time"

type BaseModel struct {
	Id        int        `json:"-" gorm:"column:id;" db:"id"`
	Uid       *UID       `json:"id" gorm:"-"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"  db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"  db:"updated_at"`
}

func NewBaseModel() BaseModel {
	now := time.Now().UTC()

	return BaseModel{
		Id:        0,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}

func (model *BaseModel) Mask(objectId int) {
	uid := NewUID(uint32(model.Id), objectId, 1)
	model.Uid = &uid
}
