package utils

import (
	"gorm.io/gorm"
	"math"
)

const (
	DefaultSize int = 8
	DefaultPageOffset int = 0
)

type Pagination struct {
	Limit        int         `json:"limit,omitempty;query:limit"`
	Page         int         `json:"page,omitempty;query:page"`
	Sort         string      `json:"sort,omitempty;query:sort"`
}

type PaginationResponse struct {
	TotalRows  		int         `json:"total_rows"`
	TotalPages 		int         `json:"total_pages"`
	CurrentPage  	int      		`json:"current_page"`
}

type PaymentResponse struct {
	PaginationResponse
	Rows      interface{} `json:"payments"`
}

type TransactionResponse struct {
	PaginationResponse
	Rows interface{} `json:"transactions"`
}

func (p *Pagination) Paginate(page, size int, order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		p.Page = page
		p.Limit = getLimit(size)
		p.Sort = order
		return db.Offset(getOffset(p.Page, p.Limit)).Limit(p.Limit).Order(p.Sort)
	}
}

func SetResponse(p *Pagination, model string, rows interface{}, count int64) interface{} {
	totalRows := int(count)
	totalPages := int(math.Ceil(float64(totalRows / p.Limit)))
	if totalPages == 0 {
		totalPages = 1
	}
	currentPage := p.Page
	paginationResponse := PaginationResponse{ totalRows, totalPages, currentPage }

	var r interface{}
	if model == "payments" {
		r = PaymentResponse{paginationResponse, rows}
	} else if model == "transactions" {
		r = TransactionResponse{paginationResponse, rows}
	}
	return r
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
