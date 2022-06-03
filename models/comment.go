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

func (m *Models) UpdateComment(photoId int, message string) (*Comment, error) {
	var c Comment

	q := `
	UPDATE Comments
	SET message = $1,
	WHERE id = $2
	RETURNING id, message, user_id, photo_id`

	stmt, err := m.db.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(message, photoId).Scan(&c.Id, &c.Message, &c.UserId, &c.PhotoId)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (m *Models) DeleteComment(commentId int) (int, error) {
	q := `
	DELETE FROM Comments
	WHERE id = $1
	RETURNING id`

	stmt, err := m.db.Prepare(q)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	var id int

	err = stmt.QueryRow(commentId).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
