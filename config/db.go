package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "go-memo-api/models"
)

func InitDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("notes.db"), &gorm.Config{}) // ← 直接 `sqlite.Open` を使う
    if err != nil {
        log.Fatal("Failed to connect to database")
    }
    db.AutoMigrate(&models.Note{})
    return db
}
