package dto

type PagePagination struct {
	Page int
	Next int
	Prev int
	Data []FilDto
}
