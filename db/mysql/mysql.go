package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/dgkg/keypass/db/sqlite"
	"github.com/dgkg/keypass/model"
)

type DBMysql = sqlite.SQLite

func New() *DBMysql {
	var conn DBMysql

	dsn := "root:example@tcp(127.0.0.1:3306)/keypass?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})
	conn.SetDB(db)

	return &conn
}
