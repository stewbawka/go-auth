package controllers

import (
    "net/http"
    "reflect"
    "time"
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

    token := models.Token{UserID: user.ID}
    database.DBConn.Create(&token)

    c.SetCookie("token", token.Token, 60*60*24, "/", "localhost", true, true)
    c.JSON(http.StatusCreated, gin.H{"data": token})
}

func InvalidateTokenByCookie(c *gin.Context) {
    tokenString, err := c.Cookie("token")
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Token unauthorized!"})
        return
    }
    var token models.Token
    if err := database.DBConn.Where("token = ?", tokenString).Joins("User").First(&token).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Token unauthorized!"})
        return
    }

    database.DBConn.Model(&token).Updates(models.Token{InvalidatedAt: time.Now() })

    // TODO: can clear it without specifying a TTL?
    c.SetCookie("token", "", 1, "/", "localhost", true, true)
    c.Status(http.StatusNoContent)
}

func GetTokenByCookie(c *gin.Context) {
    tokenString, err := c.Cookie("token")
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Token unauthorized!"})
        return
    }
    var token models.Token
    if err := database.DBConn.Where("token = ?", tokenString).Where("invalidated_at is null").Joins("User").First(&token).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Token unauthorized!"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": token})
}
