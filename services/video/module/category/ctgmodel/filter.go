package ctgmodel

import "github.com/caovanhoang63/hiholive/shared/golang/core"

type CategoryFilter struct {
	core.BaseFilter `json:",inline"`
	Name            string `json:"name,omitempty" form:"name,omitempty"`
}
