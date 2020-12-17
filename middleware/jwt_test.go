package middleware

import "testing"

func TestNewJWT(t *testing.T) {
	jwtValue, err := NewJWT("03d924ed9af94311848eb31081a5cb59", "Bob L'éponge")
	if err != nil {
		t.Errorf("try to create JWT value %v", err)
	}
	t.Log("Success creating JWT value", jwtValue)
}
