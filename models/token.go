package models

import (
    "time"
)

type Token struct {
    ID     uint   `json:"id" gorm:"primary_key"`
    UserID int `json:"user_id"`
    User User
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
