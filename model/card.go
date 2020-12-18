package model

import (
	"time"
)

type Card struct {
	ID string `json:"uuid" db:"id"`
	// UserID refer to the User.ID
	UserID              string `json:"user_id" db:"user_id"`
	URL                 string `json:"url" db:"url"`
	UserAccountLogin    string `json:"user_account_login" db:"user_account_login"`
	UserAccountPassword string `json:"user_account_password" db:"user_account_password"`
	// Activated the cart should be activated before usage.
	Activated bool `json:"activated" db:"activated"`
	// Pic is the uri of the image source link.
	Pic          string    `json:"pic" db:"pic"`
	CreationDate time.Time `json:"creation_date" db:"creation_date"`
}

func (c *Card) ValidatePayload() []error {
	return nil
}
