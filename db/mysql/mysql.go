package mysql

import (
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"

	"github.com/dgkg/keypass/db"
	"github.com/dgkg/keypass/model"
)

var schema string

func init() {
	data, err := ioutil.ReadFile("./schema.sql")
	if err != nil {
		panic(err)
	}
	schema = string(data)
}

var _ db.DB = &DBMysql{}

type DBMysql struct {
	conn *sqlx.DB
}

func New(dsn string) *DBMysql {
	log.Println("mysql: trys to connect to database")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("mysql: PING", db.Ping())
	log.Println(db.MustExec(schema).RowsAffected())

	return &DBMysql{
		conn: db,
	}
}

func (db *DBMysql) CreateUser(u *model.User) (*model.User, error) {
	u.ID = uuid.NewV4().String()
	createUser := `INSERT INTO users (id, first_name, last_name, email, password, creation_date) VALUES (?, ?, ?, ?, ? , NOW());`
	db.conn.MustExec(createUser, u.ID, u.FirstName, u.LastName, u.Email, u.Password)
	return u, nil
}

func (db *DBMysql) GetUser(uuid string) (*model.User, error) {
	var u model.User
	return &u, db.conn.Get(&u, "SELECT * FROM users WHERE id = ?", uuid)
}

func (db *DBMysql) GetUserByEmail(email string) (*model.User, error) {
	var u model.User
	return &u, db.conn.Get(&u, "SELECT * FROM users WHERE email = ?", email)
}

func (db *DBMysql) DeleteUser(uuid string) error {
	r, err := db.conn.Query("DELETE FROM users WHERE id = ?;", uuid)
	if err != nil {
		return err
	}
	return r.Err()
}

func (db *DBMysql) UpdateUser(uuid string, payload *model.Payloadpatch) (*model.User, error) {
	return nil, nil
}

func (db *DBMysql) GetAllUser() ([]*model.User, error) {
	var us []*model.User
	return us, db.conn.Select(&us, "SELECT * FROM users")
}

func (db *DBMysql) CreateCard(u *model.Card) (*model.Card, error) {
	return nil, nil
}

func (db *DBMysql) GetCard(uuid string) (*model.Card, error) {
	return nil, nil
}

func (db *DBMysql) DeleteCard(uuid string) error {
	return nil
}

func (db *DBMysql) UpdateCard(uuid string, payload *model.Payloadpatch) (*model.Card, error) {
	return nil, nil
}

func (db *DBMysql) GetAllCard() ([]*model.Card, error) {
	return nil, nil
}
