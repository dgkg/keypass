package sqlite

import (
	"github.com/dgkg/keypass/db"
	"github.com/dgkg/keypass/model"
)

var _ db.DB = &SQLite{}

type SQLite struct{}

func New() *SQLite {
	return nil
}

func (db *SQLite) CreateUser(u *model.User) (*model.User, error) {
	return nil, nil
}
func (db *SQLite) GetUser(uuid string) (*model.User, error) {
	return nil, nil
}
func (db *SQLite) DeleteUser(uuid string) (*model.User, error) {
	return nil, nil
}
func (db *SQLite) UpdateUser(uuid string, payload *model.Payloadpatch) (*model.User, error) {
	return nil, nil
}
func (db *SQLite) GetAllUser() ([]*model.User, error) {
	return nil, nil
}
