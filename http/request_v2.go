package http

// ListOptionsV2 list options struct
type ListOptionsV2 struct {
	Page     int  `form:"page"`
	PageSize int  `form:"page_size"`
	Count    bool `form:"count"`
}

// IsPagination is pagination
func (o *ListOptionsV2) IsPagination() bool {
	return o.Page > 0 && o.PageSize > 0
}

// Limit return limit for sql query
func (o *ListOptionsV2) Limit() int {
	return o.PageSize
}

// Offset return offset for sql query
func (o *ListOptionsV2) Offset() int {
	return (o.Page - 1) * o.PageSize
}
