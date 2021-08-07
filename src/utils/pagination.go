package utils

import (
	"gorm.io/gorm"
)

const (
	DefaultSize int = 8
	DefaultPageOffset int = 0
)

type Pagination struct {
	Limit        int         `json:"limit,omitempty;query:limit"`
	Page         int         `json:"page,omitempty;query:page"`
	Sort         string      `json:"sort,omitempty;query:sort"`
	Rows         interface{} `json:"rows"`
}

type PaginationResponse struct {
	TotalRows int `json:"total_rows"`
	Payments      interface{} `json:"payments"`
	TotalPages   int				 `json:"total_pages"`
	CurrentPage  int      `json:"current_page"`
}

func (p *Pagination) Paginate(db *gorm.DB, page, size int, order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		p.Page = page
		p.Limit = getLimit(size)
		p.Sort = order
		return db.Offset(getOffset(p.Page, p.Limit)).Limit(p.Limit).Order(p.Sort)
	}
}

func getLimit(size int) int {
	if size == 0 {
		return DefaultSize
	}
	return +size
}

func getOffset(page, limit int) int {
	if page == 0 {
		return DefaultPageOffset
	}
	return page * limit
}
