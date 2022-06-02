package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"hacktiv8.com/assignment-4/routers"
)

func main() {
	l := log.New(os.Stdout, "backend-api", log.LstdFlags)
	r := routers.NewRoute(l)

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
