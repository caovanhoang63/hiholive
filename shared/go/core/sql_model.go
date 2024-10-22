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

func (sqlModel *BaseModel) Mask(objectId int) {
	uid := NewUID(uint32(sqlModel.Id), objectId, 1)
	sqlModel.Uid = &uid
}
