package model

type Book struct {
	BookId string `json:"id" db:"book_id"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
}
