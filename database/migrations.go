package database

import (
    "github.com/stewbawka/go-auth/models"
)

func Migrate() {
    DBConn.AutoMigrate(models.User{})
    DBConn.AutoMigrate(models.Token{})
}
