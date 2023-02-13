package model

import (
	"github.com/jackc/pgx"
)

type PgModel struct {
	Conn *pgx.Conn
}

type Model interface {
	Book(string) Book
	AllBooks() ([]Book, error)
	AddBook(Book) (Book, error)
}

func (m *PgModel) Book(id string) Book {
	book, err := GetBookFromDB(m.Conn, id)

	if err != nil {
		panic(err)
	}

	return book
}

func (m *PgModel) AddBook(book Book) (Book, error) {
	err := WriteBookToDb(m.Conn, book)

	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func (m *PgModel) AllBooks() ([]Book, error) {
	books, err := GetAllBooksFromDB(m.Conn)

	if err != nil {
		return []Book{}, err
	}
	return books, nil
}
