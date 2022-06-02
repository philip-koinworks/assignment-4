package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"hacktiv8.com/assignment-4/helpers"
)

type Users struct {
	l *log.Logger
}

type UserRegisReq struct {
	Age      int
	Email    string
	Password string
	Username string
}

type UserRes struct {
	Status     string       `json:"status"`
	StatucCode int          `json:"statusCode"`
	Data       UserRegisReq `json:"data"`
}

func NewUser(l *log.Logger) *Users {
	return &Users{l}
}

func (u *Users) HandleRegister(rw http.ResponseWriter, r *http.Request) {
	var ur UserRegisReq
	u.l.Println("Handling user register")
	err := json.NewDecoder(r.Body).Decode(&ur)
	if err != nil {
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(UserRes{
		Status:     "SUCCESS",
		StatucCode: http.StatusOK,
		Data:       ur,
	})
}
