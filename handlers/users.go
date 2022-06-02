package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"

	"hacktiv8.com/assignment-4/helpers"
)

type Users struct {
	l  *log.Logger
	db *sql.DB
}

type UserRegisReq struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserLoginReq struct {
	Email    string
	Password string
}

type UserRes struct {
	StatucCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

func NewUser(l *log.Logger, db *sql.DB) *Users {
	return &Users{l, db}
}

func (u *Users) HandleRegister(rw http.ResponseWriter, r *http.Request) {
	var ur UserRegisReq

	u.l.Println("Handling user register")

	err := json.NewDecoder(r.Body).Decode(&ur)
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	stmt := `
	INSERT INTO Users (username, email, password, age)
	VALUES($1, $2, $3, $4)`

	passHash, err := bcrypt.GenerateFromPassword([]byte(ur.Password), bcrypt.MinCost)
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	_, err = u.db.Exec(stmt, ur.Username, ur.Email, passHash, ur.Age)
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(UserRes{
		StatucCode: http.StatusOK,
		Data:       "Success registering user",
	})
}

func (u *Users) HandleLogin(rw http.ResponseWriter, r *http.Request) {
	type User struct {
		id       int
		email    string
		password string
	}
	var us User
	var ul UserLoginReq

	u.l.Println("Handling user login")

	err := json.NewDecoder(r.Body).Decode(&ul)
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	stmt := `
	SELECT id, email, password FROM Users
	WHERE email = $1`

	_ = u.db.QueryRow(stmt, ul.Email).Scan(&us.id, &us.email, &us.password)
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(us.password), []byte(ul.Password))
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusUnauthorized)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       us.id,
		"email":    us.email,
		"password": us.password,
	})

	signingKey := []byte("secret")
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(UserRes{
		StatucCode: http.StatusOK,
		Data:       tokenString,
	})
}
