package moke

import (
	"time"

	"github.com/dgkg/keypass/db"
	"github.com/dgkg/keypass/model"
	uuid "github.com/satori/go.uuid"
)

var _ db.DB = &mokeDB{}

type mokeDB struct {
	users map[string]*model.User
}

func New() *mokeDB {
	var db mokeDB
	db.users = make(map[string]*model.User)
	return &db
}

func (db *mokeDB) CreateUser(u *model.User) (*model.User, error) {
	u.ID = uuid.NewV4().String()
	u.CreationDate = time.Now()
	db.users[u.ID] = u
	return u, nil
}
func (db *mokeDB) GetUser(uuid string) (*model.User, error) {

	return db.users[uuid], nil
}
func (db *mokeDB) DeleteUser(uuid string) (*model.User, error) {
	u, err := db.GetUser(uuid)
	if err != nil {
		return nil, err
	}
	delete(db.users, uuid)
	return u, nil
}
func (db *mokeDB) UpdateUser(uuid string, payload *model.Payloadpatch) (*model.User, error) {
	u, err := db.GetUser(uuid)
	if err != nil {
		return nil, err
	}

	u.FirstName = payload.ToString("first_name")
	u.LastName = payload.ToString("last_name")
	u.Email = payload.ToString("email")

	return u, nil
}
func (db *mokeDB) GetAllUser() ([]*model.User, error) {
	us := make([]*model.User, len(db.users))
	for _, u := range db.users {
		us = append(us, u)
	}
	return us, nil
}
