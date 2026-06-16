package dto

type PagePaginationFil struct {
	Query  string
	Page   int
	Next   int
	Prev   int
	NbrVis string
	Data   []FilDto
}
