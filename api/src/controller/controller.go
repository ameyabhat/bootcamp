package controller

import (
	"generate/workshop/src/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
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
		intId, err := strconv.Atoi(id)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, pg.Book(int64(intId)))
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

		insertedBook, err := pg.AddBook(book)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to add a book")
			panic(err)
			return
		}

		c.JSON(http.StatusOK, insertedBook)
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	return r
}
