package controllers

import (
    "log"
    "net/http"
    "reflect"
    "github.com/gin-gonic/gin"
    "github.com/stewbawka/go-auth/database"
    "github.com/stewbawka/go-auth/models"
	"github.com/go-playground/validator/v10"
)


func FindUsers(c *gin.Context) {
    var users []models.User
    database.DBConn.Find(&users)

    c.JSON(http.StatusOK, gin.H{"data": users})
}

func FindUser(c *gin.Context) {
    var user models.User
    if err := database.DBConn.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

type CreateUserSchema struct {
    Email string `json:"email" binding:"required"`
    FirstName string `json:"first_name" binding:"required"`
    LastName string `json:"last_name" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
    var data CreateUserSchema
	log.Print("Logging in Go!")
    if err := c.ShouldBindJSON(&data); err != nil {
        e := make(map[string]string)
        errors, _ := err.(validator.ValidationErrors)
        for _, fe := range errors {
            field, _ := reflect.TypeOf(&data).Elem().FieldByName(fe.Field())
            e[field.Tag.Get("json")] = fe.Tag()
        }
        c.JSON(http.StatusBadRequest, gin.H{"errors": e})
		return

    }

    user := models.User{Email: data.Email, FirstName: data.FirstName, LastName: data.LastName, Password: data.Password}
    database.DBConn.Create(&user)

    c.JSON(http.StatusCreated, gin.H{"data": user})
}

type UpdateUserSchema struct {
    Email string `json:"email"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
}

func UpdateUser(c *gin.Context) {
    var user models.User
    if err := database.DBConn.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
        return
    }
    var data UpdateUserSchema
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DBConn.Model(&user).Updates(models.User{Email: data.Email, FirstName: data.FirstName, LastName: data.LastName})

    c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
    var user models.User
    if err := database.DBConn.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
        return
    }

    database.DBConn.Delete(&user)

    c.Status(http.StatusNoContent)
}

