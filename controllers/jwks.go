package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/stewbawka/go-auth/jwt"
    "encoding/json"
)

func GetJwks(c *gin.Context) {
    jwksJson, err := json.Marshal(jwt.Jwks)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return 
    }

    c.Data(http.StatusOK, "application/json", jwksJson)
}
