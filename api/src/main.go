package main

import (
	"fmt"
	"generate/workshop/src/controller"
	"generate/workshop/src/model"
	"os"

	"github.com/jackc/pgx"
)

func main() {

	conn, err := pgx.Connect(pgx.ConnConfig{
		User:     "user",
		Database: "bootcamp",
		Password: "pwd",
		Host:     "localhost",
		Port:     5433,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()

	m := &model.PgModel{
		Conn: conn,
	}
	c := &controller.PgController{
		Model: m,
	}
	c.Serve().Run(":8080")
}
