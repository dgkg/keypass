package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/dgkg/keypass/db"
	"github.com/dgkg/keypass/db/sqlite"
	"github.com/dgkg/keypass/model"
)

var _ db.DB = &DBMysql{}

type DBMysql = sqlite.SQLite

func New(dsn string) *DBMysql {
	var conn DBMysql

	//dsn := "root:example@tcp(127.0.0.1:3306)/keypass?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Card{})

	conn.SetDB(db)

	return &conn
}
