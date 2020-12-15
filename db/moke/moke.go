package moke

import "github.com/dgkg/keypass/model"

type mokeDB struct {
	users map[string]*model.User
}

func New() *mokeDB {
	var db mokeDB
	db.users = make(map[string]*model.User)
	return &db
}

func (db *mokeDB) CreateUser(u *model.User) (*model.User, error) {
	return nil, nil
}
func (db *mokeDB) GetUser(uuid string) (*model.User, error) {
	return nil, nil
}
func (db *mokeDB) DeleteUser(uuid string) (*model.User, error) {
	return nil, nil
}
func (db *mokeDB) UpdateUser(uuid string, data map[string]interface{}) (*model.User, error) {
	return nil, nil
}
func (db *mokeDB) GetAllUser() ([]model.User, error) {
	return nil, nil
}
