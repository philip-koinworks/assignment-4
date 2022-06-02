package helpers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status int
	Data   string
}

func ServerError(rw http.ResponseWriter, err error, code int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusInternalServerError)
	r := response{
		Status: code,
		Data:   err.Error(),
	}
	json.NewEncoder(rw).Encode(r)
}
