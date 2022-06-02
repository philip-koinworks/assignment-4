package models

import (
	_ "github.com/lib/pq"
)

type User struct {
	Id       int
	Email    string
	Password string
}

func (m *Models) InsertUser(username, email string, pass []byte, age int) (int, error) {
	var id int

	q := `
	INSERT INTO Users (username, email, password, age)
	VALUES($1, $2, $3, $4) RETURNING id`

	stmt, err := m.db.Prepare(q)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username, email, pass, age).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (m *Models) GetOneUser(email string) (*User, error) {
	var us User

	stmt := `
	SELECT id, email, password FROM Users
	WHERE email = $1`

	_ = m.db.QueryRow(stmt, email).Scan(&us.Id, &us.Email, &us.Password)

	return &us, nil
}
