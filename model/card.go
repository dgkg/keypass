package model

import (
	"net/url"
	"time"
)

type Card struct {
	ID string
	// UserID refer to the User.ID
	UserID              string
	URL                 url.URL
	UserAccountLogin    string
	UserAccountPassword string
	CreationDate        time.Time
	// Activated the cart should be activated before usage.
	Activated bool
	// Pic is the uri of the image source link.
	Pic string
}
