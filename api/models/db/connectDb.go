package db

import (
	"app/config"
	"app/models/entity"
	"fmt"

	"github.com/jinzhu/gorm"
)

// ConnectDb DB接続
func ConnectDb() *gorm.DB {
	DBMS := config.Config.DBMS
	USER := config.Config.USER
	PASS := config.Config.PASS
	PROTOCOL := config.Config.PROTOCOL
	DBNAME := config.Config.DBNAME
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("データベースへの接続に失敗しました")
	}
	db.LogMode(true)
	db.AutoMigrate(&entity.Tweet{})
	db.AutoMigrate(&entity.User{})
	fmt.Println("db connected: ", &db)

	return db
}

// Close DB切断
func Close() {
	ConnectDb().Close()
}
