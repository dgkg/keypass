package model

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserUnmarshalJSON(t *testing.T) {
	data := []struct {
		title   string
		payload []byte
		result  User
		err     error
	}{
		{
			"A",
			[]byte(`{"first_name":"Bob", "last_name":"l'éponge"}`),
			User{FirstName: "Bob", LastName: "l'éponge"},
			nil,
		},
		{
			"B",
			[]byte(`{"first_name":1, "last_name":"l'éponge"}`),
			User{FirstName: "Bob", LastName: "l'éponge"},
			errors.New("json: cannot unmarshal number into Go struct field .first_name of type string"),
		},
	}

	for _, d := range data {
		var u User
		if err := json.Unmarshal(d.payload, &u); err != nil {
			if err.Error() == d.err.Error() {
				continue
			}
			t.Errorf("try to map payload data with unmarshal %v", err)
			continue
		}
		assert.Equal(t, u.FirstName, d.result.FirstName, "they should have the same first name.")
		assert.Equal(t, u.LastName, d.result.LastName, "they should have the same last name.")
	}
}
