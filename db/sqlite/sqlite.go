package sqlite

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/dgkg/keypass/db"
	"github.com/dgkg/keypass/model"
	uuid "github.com/satori/go.uuid"
)

var _ db.DB = &SQLite{}

type SQLite struct {
	db *gorm.DB
}

func New(dbName string) *SQLite {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Card{})

	return &SQLite{
		db: db,
	}
}

func (db *SQLite) SetDB(dbgorm *gorm.DB) {
	db.db = dbgorm
}

func (db *SQLite) CreateUser(u *model.User) (*model.User, error) {
	u.ID = uuid.NewV4().String()
	u.CreationDate = time.Now()
	db.db.Create(&u)
	return u, nil
}

func (db *SQLite) GetUser(uuid string) (*model.User, error) {
	var u model.User
	db.db.Where("id = ?", uuid).First(&u)
	return &u, nil
}

func (db *SQLite) DeleteUser(uuid string) (*model.User, error) {
	var u model.User
	db.db.Where("id = ?", uuid).Delete(&u)
	return &u, nil
}

func (db *SQLite) UpdateUser(uuid string, payload *model.Payloadpatch) (*model.User, error) {

	db.db.Model(&model.User{}).Where("id = ?", uuid).Updates(payload.Data)
	return db.GetUser(uuid)
}

func (db *SQLite) GetAllUser() ([]*model.User, error) {
	var us []*model.User
	db.db.Find(&us)
	return us, nil
}

func (db *SQLite) CreateCard(c *model.Card) (*model.Card, error) {
	c.ID = uuid.NewV4().String()
	c.CreationDate = time.Now()
	db.db.Create(&c)
	return c, nil
}

func (db *SQLite) GetCard(uuid string) (*model.Card, error) {
	var c model.Card
	db.db.Where("id = ?", uuid).First(&c)
	return &c, nil
}

func (db *SQLite) DeleteCard(uuid string) (*model.Card, error) {
	var u model.Card
	db.db.Where("id = ?", uuid).Delete(&u)
	return &u, nil
}

func (db *SQLite) UpdateCard(uuid string, payload *model.Payloadpatch) (*model.Card, error) {
	db.db.Model(&model.Card{}).Where("id = ?", uuid).Updates(payload.Data)
	return db.GetCard(uuid)
}

func (db *SQLite) GetAllCard() ([]*model.Card, error) {
	var cs []*model.Card
	db.db.Find(&cs)
	return cs, nil
}
