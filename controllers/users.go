package controllers

import (
    "fmt"
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
    Password string `json:"password"`
}

func CreateUser(c *gin.Context) {
    var data CreateUserSchema
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
    FirstName *string `json:"first_name" binding:"omitempty,min=1"`
    LastName *string `json:"last_name" binding:"omitempty,min=1"`
}

func UpdateUser(c *gin.Context) {
    fmt.Println("*************")
    fmt.Println(c.Request.Header["Authorization"])
    fmt.Println("*************")
    fmt.Println("*************")
    fmt.Println(c.Request.Header["x-jwt"])
    fmt.Println("*************")
    var user models.User
    if err := database.DBConn.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
        return
    }
    var data UpdateUserSchema
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

    var updateUser models.User
    if (data.FirstName != nil) {
        updateUser.FirstName = *data.FirstName
    }
    if (data.LastName != nil) {
        updateUser.LastName = *data.LastName
    }
    database.DBConn.Model(&user).Updates(updateUser)

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

