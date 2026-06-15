package dto

type PagePaginationFil struct {
	Page   int
	Next   int
	Prev   int
	NbrVis string
	Data   []FilDto
}
