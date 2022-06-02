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
	"hacktiv8.com/assignment-4/models"
)

type Users struct {
	l  *log.Logger
	db *sql.DB
}

type UserRegisReq struct {
	Id       int    `json:"id"`
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
	StatucCode int                    `json:"statusCode"`
	Data       map[string]interface{} `json:"data"`
}

func NewUser(l *log.Logger, db *sql.DB) *Users {
	return &Users{l, db}
}

func (u *Users) HandleRegister(rw http.ResponseWriter, r *http.Request) {
	var ur UserRegisReq

	u.l.Println("Handling user register")

	um := models.NewModels(u.db)

	err := json.NewDecoder(r.Body).Decode(&ur)
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(ur.Password), bcrypt.MinCost)
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	insertedId, err := um.InsertUser(ur.Username, ur.Email, passHash, ur.Age)
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(UserRes{
		StatucCode: http.StatusOK,
		Data: map[string]interface{}{
			"email":    ur.Email,
			"age":      ur.Age,
			"username": ur.Username,
			"id":       insertedId,
		},
	})
}

func (u *Users) HandleLogin(rw http.ResponseWriter, r *http.Request) {
	var ul UserLoginReq

	u.l.Println("Handling user login")

	um := models.NewModels(u.db)

	err := json.NewDecoder(r.Body).Decode(&ul)
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	res, err := um.GetOneUser(ul.Email)
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusUnauthorized)
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(ul.Password))
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusUnauthorized)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       res.Id,
		"email":    res.Email,
		"password": res.Password,
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
		Data: map[string]interface{}{
			"token": tokenString,
		},
	})
}

func (u *Users) HandleUpdate(rw http.ResponseWriter, r *http.Request) {

}
