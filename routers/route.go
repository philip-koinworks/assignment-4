package routers

import (
	"database/sql"
	"log"

	"github.com/gorilla/mux"

	"hacktiv8.com/assignment-4/handlers"
	"hacktiv8.com/assignment-4/middlewares"
)

type route struct {
	l  *log.Logger
	db *sql.DB
}

func NewRoute(l *log.Logger, db *sql.DB) *route {
	return &route{l, db}
}

func (r *route) Route() *mux.Router {
	u := handlers.NewUser(r.l, r.db)
	p := handlers.NewPhoto(r.l, r.db)
	c := handlers.NewComments(r.l, r.db)

	rs := mux.NewRouter()

	rs.HandleFunc("/users/register", u.HandleRegister).Methods("POST")
	rs.HandleFunc("/users/login", u.HandleLogin).Methods("POST")
	rs.HandleFunc("/users/{userId:[0-9]+}", middlewares.Authenticate(u.HandleUpdate)).Methods("PUT")
	rs.HandleFunc("/users", middlewares.Authorize(u.HandleDelete)).Methods("DELETE")

	rs.HandleFunc("/photos", middlewares.Authorize(p.AddPhoto)).Methods("POST")
	rs.HandleFunc("/photos", middlewares.Authorize(p.GetPhoto)).Methods("GET")
	rs.HandleFunc("/photos/{photoId:[0-9]+}", middlewares.Authorize(p.UpdatePhoto)).Methods("PUT")
	rs.HandleFunc("/photos/{photoId:[0-9]+}", middlewares.Authorize(p.DeletePhoto)).Methods("DELETE")

	rs.HandleFunc("/comments", middlewares.Authenticate(c.AddComments)).Methods("POST")
	// rs.HandleFunc("/comments").Methods("GET")
	rs.HandleFunc("/comments/{commentId:[0-9]+}", middlewares.Authorize(c.UpdateComment)).Methods("PUT")
	// rs.HandleFunc("/comments/{commentId}").Methods("DELETE")

	// rs.HandleFunc("/socialmedias").Methods("POST")
	// rs.HandleFunc("/socialmedias").Methods("GET")
	// rs.HandleFunc("/socialmedias/{socialMediaId}").Methods("PUT")
	// rs.HandleFunc("/socialmedias/{socialMediaId}").Methods("DELETE")
	return rs
}
