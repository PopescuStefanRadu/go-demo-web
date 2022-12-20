package query

type PaginatedResponse[T any] struct {
	Data        []T
	TotalPages  int
	CurrentPage int
	PageSize    int
}

type Paged struct {
	Number *int
	Size   *int
}

type Sorted struct {
	SortBy []Sort
}

type Sort struct {
	FieldName string
	Order     Order
}

type Order string

const (
	OrderAscending Order = "ASC"
	OrderDescending Order = "DESC"
)
