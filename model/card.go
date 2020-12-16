package model

import (
	"time"
)

type Card struct {
	ID string
	// UserID refer to the User.ID
	UserID              string
	URL                 string
	UserAccountLogin    string
	UserAccountPassword string
	CreationDate        time.Time
	// Activated the cart should be activated before usage.
	Activated bool
	// Pic is the uri of the image source link.
	Pic string
}

func (c *Card) ValidatePayload() []error {
	return nil
}
