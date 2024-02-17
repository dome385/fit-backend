package main

import (
	"fit-backend/internal/repository"
	"fit-backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

const port = 8080

func main() {

	var app application

	flag.StringVar(&app.DSN, "dsn", "user=postgres dbname=postgres password=docker host=localhost port=5432 sslmode=disable", "Postgres connection string")

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
