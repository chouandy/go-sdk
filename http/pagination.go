package http

// Pagination pageable struct
type Pagination struct {
	Page      int64 `form:"page"`
	PageSize  int64 `form:"page_size"`
	TotalSize int64
}

// IsPagination is pagination
func (p *Pagination) IsPagination() bool {
	return p.Page > 0 && p.PageSize > 0
}

// Limit return limit for sql query
func (p *Pagination) Limit() int64 {
	return p.PageSize
}

// Offset return offset for sql query
func (p *Pagination) Offset() int64 {
	return (p.Page - 1) * p.PageSize
}

// Pages return pages
func (p *Pagination) Pages() int64 {
	pages := p.TotalSize / p.PageSize
	if p.TotalSize%p.PageSize > 0 {
		pages++
	}

	return pages
}
