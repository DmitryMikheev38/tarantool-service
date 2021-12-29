package dto

type GetListAuthorsParams struct {
	BooksLimit int    `form:"booksLimit"`
	Limit      uint32 `form:"limit"`
	Offset     uint32 `form:"offset"`
}
