package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func ConnectDatabase() (*gorm.DB, error) {
    database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    if err := database.AutoMigrate(&Reservation{}); err != nil {
        return nil, err
    }
    return database, nil
}