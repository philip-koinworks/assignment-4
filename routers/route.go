package routers

import (
	"log"

	"github.com/gorilla/mux"

	"hacktiv8.com/assignment-4/handlers"
)

type route struct {
	l *log.Logger
}

func NewRoute(l *log.Logger) *route {
	return &route{l}
}

func (r *route) Route() *mux.Router {
	u := handlers.NewUser(r.l)
	rs := mux.NewRouter()
	rs.HandleFunc("/users/register", u.HandleRegister).Methods("POST")
	// rs.HandleFunc("/users/login").Methods("POST")
	// rs.HandleFunc("/users").Methods("PUT")
	// rs.HandleFunc("/users").Methods("DELETE")

	// rs.HandleFunc("/photos").Methods("POST")
	// rs.HandleFunc("/photos").Methods("GET")
	// rs.HandleFunc("/photos/{photoId}").Methods("PUT")
	// rs.HandleFunc("/photos/{photoId}").Methods("DELETE")

	// rs.HandleFunc("/comments").Methods("POST")
	// rs.HandleFunc("/comments").Methods("GET")
	// rs.HandleFunc("/comments/{commentId}").Methods("PUT")
	// rs.HandleFunc("/comments/{commentId}").Methods("DELETE")

	// rs.HandleFunc("/socialmedias").Methods("POST")
	// rs.HandleFunc("/socialmedias").Methods("GET")
	// rs.HandleFunc("/socialmedias/{socialMediaId}").Methods("PUT")
	// rs.HandleFunc("/socialmedias/{socialMediaId}").Methods("DELETE")
	return rs
}
