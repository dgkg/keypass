package moke

import (
	"time"

	database "github.com/dgkg/keypass/db"
	"github.com/dgkg/keypass/model"
	uuid "github.com/satori/go.uuid"
)

var _ database.DB = &mokeDB{}

type mokeDB struct {
	users     map[string]*model.User
	cards     map[string]*model.Card
	conteners map[string]*model.Contener
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

func (db *mokeDB) GetUserByEmail(email string) (*model.User, error) {
	for _, u := range db.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, database.NewErrNotFound("email"+email, nil)
}

func (db *mokeDB) DeleteUser(uuid string) error {
	delete(db.users, uuid)
	return nil
}
func (db *mokeDB) UpdateUser(uuid string, payload *model.Payloadpatch) (*model.User, error) {
	u, err := db.GetUser(uuid)
	if err != nil {
		return nil, err
	}

	payload.ToString(&u.FirstName, "first_name")
	payload.ToString(&u.LastName, "last_name")
	payload.ToString(&u.Email, "email")

	return u, nil
}
func (db *mokeDB) GetAllUser() ([]*model.User, error) {
	us := make([]*model.User, len(db.users))
	for _, u := range db.users {
		us = append(us, u)
	}
	return us, nil
}

func (db *mokeDB) CreateCard(c *model.Card) (*model.Card, error) {
	c.ID = uuid.NewV4().String()
	c.CreationDate = time.Now()
	db.cards[c.ID] = c
	return c, nil
}

func (db *mokeDB) GetCard(uuid string) (*model.Card, error) {
	return db.cards[uuid], nil
}

func (db *mokeDB) DeleteCard(uuid string) error {
	delete(db.cards, uuid)
	return nil
}

func (db *mokeDB) UpdateCard(uuid string, payload *model.Payloadpatch) (*model.Card, error) {
	// TODO implement this function.
	return nil, nil
}

func (db *mokeDB) GetAllCard() ([]*model.Card, error) {
	us := make([]*model.Card, len(db.cards))
	for _, u := range db.cards {
		us = append(us, u)
	}
	return us, nil
}

func (db *mokeDB) CreateContener(c *model.Contener) (*model.Contener, error) {
	c.ID = uuid.NewV4().String()
	c.CreationDate = time.Now()
	db.conteners[c.ID] = c
	return c, nil
}

func (db *mokeDB) GetContener(uuid string) (*model.Contener, error) {
	return db.conteners[uuid], nil
}

func (db *mokeDB) DeleteContener(uuid string) error {
	delete(db.conteners, uuid)
	return nil
}

func (db *mokeDB) UpdateContener(uuid string, payload *model.Payloadpatch) (*model.Contener, error) {
	// TODO implement this function.
	return nil, nil
}

func (db *mokeDB) GetAllContener() ([]*model.Contener, error) {
	us := make([]*model.Contener, len(db.conteners))
	for _, u := range db.conteners {
		us = append(us, u)
	}
	return us, nil
}
