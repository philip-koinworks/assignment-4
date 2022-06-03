package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"hacktiv8.com/assignment-4/helpers"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if len(auth) == 0 {
			helpers.ServerError(rw, errors.New("Unauthorize"), http.StatusUnauthorized)
			return
		}
		authSplit := strings.Split(auth, " ")

		token := authSplit[1]

		t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte("secret"), nil
		})

		if err != nil {
			helpers.ServerError(rw, errors.New("Unauthorize"), http.StatusUnauthorized)
			return
		}

		if _, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
			next(rw, r)
		} else {
			helpers.ServerError(rw, errors.New("Unauthorize"), http.StatusUnauthorized)
			return
		}
	})
}
