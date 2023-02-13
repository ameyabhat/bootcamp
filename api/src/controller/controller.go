package controller

import (
	"generate/workshop/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Serve() *gin.Engine
}

type PgController struct {
	model.Model
}

// Everything above here is going to move to a  folder (controller layer)
func (pg *PgController) Serve() *gin.Engine {
	r := gin.Default()
	r.GET("/v1/books/:bookId", func(c *gin.Context) {
		id := c.Param("bookId")

		c.JSON(http.StatusOK, pg.Book(id))
	})
	r.GET("/v1/books/", func(c *gin.Context) {
		books, err := pg.AllBooks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
		}
		c.JSON(http.StatusOK, books)
	})

	r.POST("/v1/addBook", func(c *gin.Context) {
		var book model.Book

		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal book")
			return
		}

		_, err := pg.AddBook(book)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to add a book")
			panic(err)
			return
		}

		c.JSON(http.StatusOK, book.BookId)
	})

	return r
}
