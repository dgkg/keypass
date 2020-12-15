package model

import (
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
