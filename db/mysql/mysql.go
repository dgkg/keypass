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

	log.Println("mysql: wait to connect to database")
	// time.Sleep(time.Second * 60)
	log.Println("mysql: try to connect to database")

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec("SELECT VERSION();", nil)
	log.Println("mysql: select version", res, err)

	res, err = db.Exec("SHOW GRANTS;", nil)
	log.Println("mysql: show grants", res, err)

	res, err = db.Exec("SELECT 1;", nil)
	log.Println("mysql: select 1", res, err)

	log.Println("mysql: PING", db.Ping())
	// log.Println(db.MustExec(schema).RowsAffected())
	res, err = db.Exec(schemaConteners, nil)
	log.Println("mysql: create schema conteners", res, err)

	res, err = db.Exec(schemaCards, nil)
	log.Println("mysql: create schema cards", res, err)

	res, err = db.Exec(schemaUser, nil)
	log.Println("mysql: create schema users", res, err)

	return &DBMysql{
		conn: db,
	}
}

var schemaConteners string = `
CREATE TABLE conteners (
	id text NULL,
	user_id text NULL,
	title text NULL,
	secret text NULL,
	creation_date datetime NULL
 );`

var schemaUser string = `
 CREATE TABLE users (
	id text NULL,
	first_name text NULL,
	last_name text NULL,
	email text NULL,
	password text NULL,
	creation_date datetime NULL
 );`

var schemaCards string = `
 CREATE TABLE cards (
	id text NULL,
	contener_id text NULL,
	url text,
	pic blob,
	activated tinyint(1) NULL,
	creation_date datetime NULL
 );`

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

func (db *DBMysql) CreateContener(u *model.Contener) (*model.Contener, error) {
	return nil, nil
}

func (db *DBMysql) GetContener(uuid string) (*model.Contener, error) {
	return nil, nil
}

func (db *DBMysql) DeleteContener(uuid string) error {
	return nil
}

func (db *DBMysql) UpdateContener(uuid string, payload *model.Payloadpatch) (*model.Contener, error) {
	return nil, nil
}

func (db *DBMysql) GetAllContener() ([]*model.Contener, error) {
	return nil, nil
}
