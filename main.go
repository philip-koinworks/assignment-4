package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"hacktiv8.com/assignment-4/routers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "assignment4"
)

func main() {
	l := log.New(os.Stdout, "backend-api", log.LstdFlags)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	l.Println("Successfully connected to database!")
	r := routers.NewRoute(l, db)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r.Route(),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
