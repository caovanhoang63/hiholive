package core

import (
	"strings"
)

type Paging struct {
	Limit int   `json:"limit" form:"limit"`
	Page  int   `json:"page" form:"page"`
	Total int64 `json:"total" form:"-"`

	// Support cursor with UID
	FakeCursor string `json:"cursor" form:"cursor"`
	NextCursor string `json:"next_cursor"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}

	if p.Limit >= 200 {
		p.Limit = 200
	}

	p.FakeCursor = strings.TrimSpace(p.FakeCursor)
}

func (p *Paging) GetOffSet() int {
	return (p.Page - 1) * p.Limit
}
