package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

var DB *gorm.DB

func DbConnect() {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_CON"))

	if err != nil {
		fmt.Println(err)
		panic("db error")
	}

	db.DB().SetMaxIdleConns(20)
	db.DB().SetConnMaxLifetime(30 * time.Second)
	db.DB().SetMaxOpenConns(20)
	db.LogMode(true)
	DB = db

	Migrate()
}
