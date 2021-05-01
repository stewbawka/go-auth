package controllers

import (
    "net/http"
    "reflect"
    "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
    "github.com/stewbawka/go-auth/database"
    "github.com/stewbawka/go-auth/models"
)

type CreateTokenSchema struct {
    Email string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func CreateToken(c *gin.Context) {
    var data CreateTokenSchema
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

    var user models.User
    if err := database.DBConn.Where("email = ?", data.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Login Failed!"})
        return
    }

    if match, err := user.ComparePasswordAndHash(data.Password); err != nil || !match {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Login Failed!"})
        return
    }

    token := models.Token{User: user}
    database.DBConn.Create(&token)

    c.JSON(http.StatusCreated, gin.H{"data": token})
}
