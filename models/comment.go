package models

import (
	_ "github.com/lib/pq"
)

type Comment struct {
	Id      int
	Message string
	PhotoId int
	UserId  int
}

func (m *Models) InsertComment(userId float64, photoId int, message string) (*Comment, error) {
	var c Comment

	q := `
	INSERT INTO Comments (user_id, message, photo_id)
	VALUES($1, $2, $3) 
	RETURNING id, message, photo_id, user_id`

	stmt, err := m.db.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(userId, message, photoId).Scan(&c.Id, &c.Message, &c.PhotoId, &c.UserId)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
