package models

import (
	_ "github.com/lib/pq"
)

type Photo struct {
	Id       int
	Title    string
	Caption  string
	PhotoUrl string
	UserId   int
}

func (m *Models) InsertPhoto(userId float64, title, caption, url string) (*Photo, error) {
	var p Photo

	q := `
	INSERT INTO Photos (title, caption, photo_url, user_id)
	VALUES($1, $2, $3, $4) 
	RETURNING id, title, caption, photo_url, user_id`

	stmt, err := m.db.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(title, caption, url, userId).Scan(&p.Id, &p.Title, &p.Caption, &p.PhotoUrl, &p.UserId)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
