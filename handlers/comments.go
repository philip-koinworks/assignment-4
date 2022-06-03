package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"hacktiv8.com/assignment-4/helpers"
	"hacktiv8.com/assignment-4/models"
)

type Comments struct {
	l  *log.Logger
	db *sql.DB
}

type CommentReq struct {
	Id      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	PhotoId int    `json:"photo_id,omitempty"`
	UserId  int    `json:"user_id,omitempty"`
}

func NewComments(l *log.Logger, db *sql.DB) *Comments {
	return &Comments{l, db}
}

func (p *Comments) AddComments(rw http.ResponseWriter, r *http.Request) {
	var cr CommentReq
	p.l.Println("Handling add photo")

	pm := models.NewModels(p.db)
	id := r.Context().Value("id").(float64)

	err := json.NewDecoder(r.Body).Decode(&cr)
	if err != nil {
		p.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	row, err := pm.InsertComment(id, cr.PhotoId, cr.Message)
	if err != nil {
		p.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(PhotoRes{
		StatucCode: http.StatusOK,
		Data: map[string]interface{}{
			"id":       row.Id,
			"message":  row.Message,
			"photo_id": row.PhotoId,
			"user_id":  row.UserId,
		},
	})
}
