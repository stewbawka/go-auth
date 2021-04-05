package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/stewbawka/go-auth/database"
    "github.com/stewbawka/go-auth/models"
)


func FindUsers(c *gin.Context) {
    var users []models.User
    database.DBConn.Find(&users)

    c.JSON(http.StatusOK, gin.H{"data": users})
}
