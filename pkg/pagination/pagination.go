package pagination

type Pagination struct {
	Offset int
	Limit  int
}

type PagedResult[T any] struct {
	Items      []T
	TotalCount int
	HasMore    bool
}
