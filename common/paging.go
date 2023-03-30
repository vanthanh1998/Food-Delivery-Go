package common

import "strings"

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"total"` // tổng SL tìm thấy dưới DB là bao nhiêu
	// Support cursor with UID
	FaceCursor string `json:"cursor" form:"cursor"` // tối ưu tốc độ paging
	NextCursor string `json:"next_cursor"`          // tối ưu tốc độ paging
}

func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 50
	}

	p.FaceCursor = strings.TrimSpace(p.FaceCursor)
}
