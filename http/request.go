package http

// ListOptions list options struct
type ListOptions struct {
	Page     int  `form:"page"`
	PageSize int  `form:"page_size"`
	Count    bool `form:"count"`
}

// IsPagination is pagination
func (o *ListOptions) IsPagination() bool {
	return o.Page > 0 && o.PageSize > 0
}

// Limit return limit for sql query
func (o *ListOptions) Limit() int {
	return o.PageSize
}

// Offset return offset for sql query
func (o *ListOptions) Offset() int {
	return (o.Page - 1) * o.PageSize
}
