package core

import (
	"time"
)

type BaseFilter struct {
	GtCreatedAt *time.Time `json:"gtCreatedAt,omitempty" form:"gtCreatedAt,omitempty"`
	GtUpdatedAt *time.Time `json:"gtUpdatedAt,omitempty" form:"gtUpdatedAt,omitempty"`
	LtCreatedAt *time.Time `json:"ltCreatedAt,omitempty" form:"ltCreatedAt,omitempty"`
	LtUpdatedAt *time.Time `json:"ltUpdatedAt,omitempty" form:"ltUpdatedAt,omitempty"`
}
