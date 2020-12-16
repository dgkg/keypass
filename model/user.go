package model

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

// User represent a single customer used for JSON comm.
type User struct {
	ID           string    `json:"uuid"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	CreationDate time.Time `json:"creation_date"`
}

func (u *User) UnmarshalJSON(b []byte) error {
	aux := struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}{}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	u.FirstName = aux.FirstName
	u.LastName = aux.LastName
	u.Email = aux.Email
	u.Password = aux.Password

	return nil
}

func (u User) MarshalJSON() ([]byte, error) {
	aux := struct {
		ID           string    `json:"uuid"`
		FirstName    string    `json:"first_name"`
		LastName     string    `json:"last_name"`
		Email        string    `json:"email"`
		CreationDate time.Time `json:"creation_date"`
	}{
		ID:           u.ID,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		CreationDate: u.CreationDate,
	}
	return json.Marshal(aux)
}

func NewUser(fn, ln, email, pass string) (*User, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return &User{
		ID:           id.String(),
		FirstName:    fn,
		LastName:     ln,
		Email:        email,
		Password:     pass,
		CreationDate: time.Now(),
	}, nil
}
