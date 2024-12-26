package streammodel

import (
	"time"
)

type StreamUpdate struct {
	Title              *string    `json:"title" gorm:"column:title"`
	Description        *string    `json:"description" gorm:"column:description"`
	IsRerun            *bool      `json:"isRerun" gorm:"column:is_rerun"`
	Notification       *string    `json:"notification" gorm:"column:notification"`
	ScheduledStartTime *time.Time `json:"scheduledStartTime" gorm:"column:scheduled_start_time"`
	State              *string    `json:"state" gorm:"column:state"`
}
