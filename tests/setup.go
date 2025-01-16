package tests

import (
    "github.com/RedColdHearted/go-demo-exam/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func SetUpDB() (*gorm.DB, error) {
    database, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    if err := database.AutoMigrate(&models.Reservation{}); err != nil {
        return nil, err
    }
    return database, nil
}

func TearDownDB(DB *gorm.DB) {
    database, err := DB.DB()
    if err != nil {
        panic("Failing to get database instance")
    }
    database.Close()
}

