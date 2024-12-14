package stmodel

import "github.com/caovanhoang63/hiholive/shared/go/core"

type Filter struct {
	core.BaseFilter `json:",inline"`
	Name            string `json:"name,omitempty" form:"name,omitempty"`
}
