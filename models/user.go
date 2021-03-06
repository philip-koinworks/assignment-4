package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	_ "github.com/lib/pq"
)

type User struct {
	Id       int    `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json"password,omitempty"`
	Username string `json:"username,omitempty"`
	Age      int    `json:"age,omitempty"`
}

func (c User) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *User) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &c)
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

func (m *Models) UpdateUser(userId int, email, username string) (*User, error) {
	var u User

	q := `
	UPDATE Users
	SET username = $1,
		email = $2
	WHERE id = $3
	RETURNING id, age, email, username`

	stmt, err := m.db.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username, email, userId).Scan(&u.Id, &u.Age, &u.Email, &u.Username)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (m *Models) DeleteUser(userId float64) (int, error) {
	q := `
	DELETE FROM Users
	WHERE id = $1
	RETURNING id`

	stmt, err := m.db.Prepare(q)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	var id int

	err = stmt.QueryRow(userId).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
