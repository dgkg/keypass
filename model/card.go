package model

import (
	"encoding/json"
	"time"
)

type Card struct {
	ID          string `json:"uuid" db:"id"`
	ContainerID string `json:"container_uuid" db:"container_id"`
	URL         string `json:"url" db:"url"`
	// Pic is the uri of the image source link.
	Pic string `json:"pic" db:"pic"`
	// Activated the cart should be activated before usage.
	Activated    bool      `json:"activated" db:"activated"`
	CreationDate time.Time `json:"creation_date" db:"creation_date"`
}

func (c *Card) UnmarshalJSON(b []byte) error {
	aux := struct {
		ContainerID string `json:"container_uuid"`
		URL         string `json:"url"`
		Pic         string `json:"pic"`
		Activated   bool   `json:"activated"`
	}{}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	c.ContainerID = aux.ContainerID
	c.URL = aux.URL
	c.Pic = aux.Pic
	c.Activated = aux.Activated

	return nil
}

func (c Card) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

func (c *Card) ValidatePayload() []error {
	return nil
}
