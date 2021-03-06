package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

var (
    DBConn *gorm.DB
)

func Connect() {
    dsn := "root:very_secure_password@tcp(mysql:3306)/go_auth_db?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err.Error())
    }
    DBConn = db
}

func Close() {
	sqlDB, err := DBConn.DB()
	if err != nil {
        panic(err.Error())
	}
	sqlDB.Close()
}
