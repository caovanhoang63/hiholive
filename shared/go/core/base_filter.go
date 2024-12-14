package core

import (
	"time"
)

type BaseFilter struct {
	GtCreatedAt *time.Time `json:"gt_created_at,omitempty" form:"gt_created_at,omitempty"`
	GtUpdatedAt *time.Time `json:"gt_updated_at,omitempty" form:"gt_updated_at,omitempty"`
	LtCreatedAt *time.Time `json:"lt_created_at,omitempty" form:"lt_created_at,omitempty"`
	LtUpdatedAt *time.Time `json:"lt_updated_at,omitempty" form:"lt_updated_at,omitempty"`
}
