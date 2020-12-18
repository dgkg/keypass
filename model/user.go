package model

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
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

func (u *User) ValidatePayload() []error {
	var errList []error
	err := validateSize(3, 200, u.FirstName)
	if err != nil {
		errList = append(errList, err)
	}

	err = validateSize(3, 200, u.LastName)
	if err != nil {
		errList = append(errList, err)
	}

	err = validateSize(3, 320, u.Email)
	if err != nil {
		errList = append(errList, err)
	}

	if !strings.Contains(u.Email, "@") {
		errList = append(errList, errors.New("no @ found for the email"))
	} else {
		emailVals := strings.Split(u.Email, "@")
		if len(emailVals) != 2 {
			errList = append(errList, errors.New("email not valid"))
		} else {
			err = validateSize(1, 64, emailVals[1])
			if err != nil {
				errList = append(errList, err)
			}
		}
	}

	validateSize(3, 200, u.Password)
	if err != nil {
		errList = append(errList, err)
	}
	return errList
}

func validateSize(min, max int, value string) error {
	if len(value) < min || len(value) > max {
		if len(value) > max {
			return fmt.Errorf("got a value too long: %v", value[:max])
		}
		return fmt.Errorf("got a value too short: %v" + value)
	}
	return nil
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
	u.Password = Hash(aux.Password)

	return nil
}

func Hash(clear string) (hashed string) {
	h := sha256.New()
	h.Write([]byte(clear))
	return fmt.Sprintf("%x", h.Sum(nil))
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

func NewUser(fn, ln, email, pass string) *User {
	return &User{
		FirstName: fn,
		LastName:  ln,
		Email:     email,
		Password:  pass,
	}
}

type UserLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (u *UserLogin) UnmarshalJSON(b []byte) error {
	aux := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	u.Login = aux.Login
	u.Password = Hash(aux.Password)

	return nil
}

func (u UserLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}
