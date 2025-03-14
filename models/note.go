package models

import "time"

type Note struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}
