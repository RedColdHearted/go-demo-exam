package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed attempt of connection to database")
	}
	err = database.AutoMigrate(&Reservation{})
	if err != nil {
		return
	}
	DB = database
}