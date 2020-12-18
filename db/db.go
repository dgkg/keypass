package db

import (
	"github.com/dgkg/keypass/model"
)

type DB interface {
	DBUser
	DBCard
}

type DBUser interface {
	CreateUser(u *model.User) (*model.User, error)
	GetUser(uuid string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	DeleteUser(uuid string) error
	UpdateUser(uuid string, payload *model.Payloadpatch) (*model.User, error)
	GetAllUser() ([]*model.User, error)
}

type DBCard interface {
	CreateCard(u *model.Card) (*model.Card, error)
	GetCard(uuid string) (*model.Card, error)
	DeleteCard(uuid string) error
	UpdateCard(uuid string, payload *model.Payloadpatch) (*model.Card, error)
	GetAllCard() ([]*model.Card, error)
}
