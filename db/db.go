package db

import "github.com/dgkg/keypass/model"

type DB interface {
	DBUser
}

type DBUser interface {
	CreateUser(u *model.User) (*model.User, error)
	GetUser(uuid string) (*model.User, error)
	DeleteUser(uuid string) (*model.User, error)
	UpdateUser(uuid string, data map[string]interface{}) (*model.User, error)
	GetAllUser() ([]model.User, error)
}
