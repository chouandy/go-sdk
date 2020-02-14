package http

// PaginationDefaultValue
const (
	DefaultPage     int64 = 1
	DefaultPageSize int64 = 10
)

// Pagination pageable struct
type Pagination struct {
	Page      int64 `form:"page"`
	PageSize  int64 `form:"page_size"`
	TotalSize int64
}

// GetPage get page
func (p *Pagination) GetPage() int64 {
	if p.Page > 0 {
		return p.Page
	}

	return DefaultPage
}

// GetPageSize get page size
func (p *Pagination) GetPageSize() int64 {
	if p.PageSize > 0 {
		return p.PageSize
	}

	return DefaultPageSize
}

// Limit return limit for sql query
func (p *Pagination) Limit() int64 {
	return p.GetPageSize()
}

// Offset return offset for sql query
func (p *Pagination) Offset() int64 {
	return (p.GetPage() - 1) * p.GetPageSize()
}

// Pages return pages
func (p *Pagination) Pages() int64 {
	pages := p.TotalSize / p.GetPageSize()
	if p.TotalSize%p.GetPageSize() > 0 {
		pages++
	}

	return pages
}
