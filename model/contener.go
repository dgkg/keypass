package model

import (
	"encoding/json"
	"time"
)

type Contener struct {
	ID           string    `json:"uuid" db:"id"`
	UserID       string    `json:"user_id" db:"user_id"`
	Title        string    `json:"title" db:"title"`
	Secret       string    `json:"secret" db:"secret"`
	Cards        []Card    `json:"cards omitempty" db:"cards"`
	CreationDate time.Time `json:"creation_date" db:"creation_date"`
}

func (c Contener) MarshalJSON() ([]byte, error) {
	aux := struct {
		ID           string    `json:"uuid"`
		UserID       string    `json:"user_uuid"`
		Title        string    `json:"title"`
		CreationDate time.Time `json:"creation_date"`
	}{
		ID:           c.ID,
		UserID:       c.UserID,
		Title:        c.Title,
		CreationDate: c.CreationDate,
	}
	return json.Marshal(aux)
}

func (c *Contener) ValidatePayload() []error {
	// TODO implement this function.
	return nil
}
