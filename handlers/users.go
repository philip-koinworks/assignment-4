package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"

	"hacktiv8.com/assignment-4/helpers"
	"hacktiv8.com/assignment-4/models"
)

type Users struct {
	l  *log.Logger
	db *sql.DB
}

type UserReq struct {
	Id       int    `json:"id,omitempty"`
	Age      int    `json:"age,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
}

type UserRes struct {
	StatucCode int                    `json:"statusCode"`
	Data       map[string]interface{} `json:"data"`
}

func NewUser(l *log.Logger, db *sql.DB) *Users {
	return &Users{l, db}
}

func (u *Users) HandleRegister(rw http.ResponseWriter, r *http.Request) {
	var ur UserReq

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
	var ul UserReq

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
	u.l.Println("Handling user update")
	var uu UserReq

	um := models.NewModels(u.db)

	err := json.NewDecoder(r.Body).Decode(&uu)
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}
	vars := mux.Vars(r)
	val, ok := vars["userId"]
	if ok != true {
		u.l.Println(err)
		helpers.ServerError(rw, errors.New("Can't find user id params"), http.StatusInternalServerError)
	}

	userId, err := strconv.Atoi(val)
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	row, err := um.UpdateUser(userId, uu.Email, uu.Username)
	if err != nil {
		u.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(UserRes{
		StatucCode: http.StatusOK,
		Data: map[string]interface{}{
			"id":       row.Id,
			"email":    row.Email,
			"username": row.Username,
			"age":      row.Age,
		},
	})
}
