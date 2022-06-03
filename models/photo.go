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
	Users    User
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

func (m *Models) SelectAllPhotos(userId float64) ([]*Photo, error) {
	var ps []*Photo

	q := `
	SELECT Photos.id, Photos.title, Photos.caption, Photos.photo_url, Photos.user_id, json_build_object('email', Users.email, 'username', Users.username) as Users
	FROM Photos
	LEFT JOIN Users
	ON Photos.user_id = $1`

	stmt, err := m.db.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var p Photo

		err := rows.Scan(&p.Id, &p.Title, &p.Caption, &p.PhotoUrl, &p.UserId, &p.Users)
		if err != nil {
			return nil, err
		}
		ps = append(ps, &p)
	}

	return ps, nil
}
