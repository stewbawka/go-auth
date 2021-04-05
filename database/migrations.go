package database

import (
    "github.com/stewbawka/go-auth/models"
)

func Migrate() {
    DBConn.AutoMigrate(models.User{})
}
