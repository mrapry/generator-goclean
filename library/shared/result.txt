package shared

import "math"

// Result common output
type Result struct {
	Data  interface{}
	Error error
}

// SliceResult include meta
type SliceResult struct {
	Data interface{}
	Meta Meta
}

// Meta model
type Meta struct {
	Page         int `json:"page"`
	Limit        int `json:"limit"`
	TotalRecords int `json:"total_records"`
	TotalPages   int `json:"total_pages"`
}

// NewMeta create new meta for slice data
func NewMeta(page, limit, totalRecords int) (m Meta) {
	m.Page, m.Limit, m.TotalRecords = page, limit, totalRecords
	m.CalculatePages()
	return
}

// CalculatePages meta method
func (m *Meta) CalculatePages() {
	if m.Limit == 0 {
		m.TotalPages = 0
	} else {
		m.TotalPages = int(math.Ceil(float64(m.TotalRecords) / float64(m.Limit)))
	}
}
