package helpers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	status int
	data   string
}

func ServerError(rw http.ResponseWriter, err error, code int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusInternalServerError)
	r := response{
		status: code,
		data:   err.Error(),
	}
	json.NewEncoder(rw).Encode(r)
}
