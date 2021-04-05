package models

import (
    "time"
    "gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
    ID     uint   `json:"id" gorm:"primary_key"`
    Email string `json:"email"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`

}
