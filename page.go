package yugle

const (
	DefaultPage = 1
	DefaultSize = 10
)

type PageParams struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

type Pagination struct {
	Page       int
	Size       int
	Total      int
	Content    interface{}
	HasPrev    bool
	HasNext    bool
	TotalPages int
}

func NewPageParams() *PageParams {
	return &PageParams{
		Page: DefaultPage,
		Size: DefaultSize,
	}
}

func GenPage(page int, size int, total int, data interface{}) *Pagination {
	pagination := &Pagination{}
	pagination.Page = page
	pagination.Size = size
	pagination.Total = total
	pagination.Content = data
	totalPages := total / size
	if total%size != 0 {
		totalPages += 1
	}
	pagination.HasNext = page < totalPages
	pagination.HasPrev = page > 1
	pagination.TotalPages = totalPages
	return pagination
}
