package model

import (
	"github.com/jackc/pgx"
)

type PgModel struct {
	Conn *pgx.Conn
}

type Model interface {
	Book(string) Book
	AddBooks(Book) (Book, error)
}

func (m *PgModel) Book(id string) Book {
	book, err := GetBooksFromDB(m.Conn, id)

	if err != nil {
		panic(err)
	}

	return book
}

func (m *PgModel) AddBooks(book Book) (Book, error) {
	err := WriteBooksToDb(m.Conn, book)

	if err != nil {
		return Book{}, err
	}

	return book, nil
}
