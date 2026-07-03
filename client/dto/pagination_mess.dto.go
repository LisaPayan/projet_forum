package dto

type PagePaginationMess struct {
	IdFil  int
	Page   int
	Next   int
	Prev   int
	NbrVis string
	Data   []MessageDto
	Tri    string
}
