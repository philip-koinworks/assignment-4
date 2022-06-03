package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"hacktiv8.com/assignment-4/helpers"
	"hacktiv8.com/assignment-4/models"
)

type Photos struct {
	l  *log.Logger
	db *sql.DB
}

type PhotoReq struct {
	Id       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Caption  string `json:"caption,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
	UserId   string `json:"user_id,omitempty"`
}

type PhotoRes struct {
	StatucCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

func NewPhoto(l *log.Logger, db *sql.DB) *Photos {
	return &Photos{l, db}
}

func (p *Photos) AddPhoto(rw http.ResponseWriter, r *http.Request) {
	var pl PhotoReq
	p.l.Println("Handling add photo")

	pm := models.NewModels(p.db)
	id := r.Context().Value("id").(float64)

	err := json.NewDecoder(r.Body).Decode(&pl)
	if err != nil {
		p.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	row, err := pm.InsertPhoto(id, pl.Title, pl.Caption, pl.PhotoUrl)
	if err != nil {
		p.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(PhotoRes{
		StatucCode: http.StatusOK,
		Data: map[string]interface{}{
			"id":        row.Id,
			"title":     row.Title,
			"caption":   row.Caption,
			"photo_url": row.PhotoUrl,
			"user_id":   row.UserId,
		},
	})
}

func (p *Photos) GetPhoto(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling get photos")

	pm := models.NewModels(p.db)
	id := r.Context().Value("id").(float64)

	rows, err := pm.SelectAllPhotos(id)
	if err != nil {
		p.l.Println(err)
		helpers.ServerError(rw, err, http.StatusInternalServerError)
	}

	res, _ := json.Marshal(rows)

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(PhotoRes{
		StatucCode: http.StatusOK,
		Data:       string(res),
	})
}
