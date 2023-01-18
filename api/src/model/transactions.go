package model

import (
	"fmt"

	"github.com/jackc/pgx"
)

func WriteBooksToDb(pool *pgx.Conn, book Book) error {
	_, err := pool.Exec(fmt.Sprintf("INSERT INTO books (book_id, title, author) VALUES (%s,%s,%s);", book.BookId, book.Title, book.Author))

	return err
}

func GetBooksFromDB(pool *pgx.Conn, book_id string) (Book, error) {
	var book Book
	err := pool.QueryRow(fmt.Sprintf("SELECT * FROM books where book_id = %s;", book_id)).Scan(&book)

	if err != nil {
		panic(err)
	}

	return book, nil
}
