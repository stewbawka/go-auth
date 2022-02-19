package main

import (
    "github.com/gin-gonic/gin"
    "github.com/stewbawka/go-auth/database"
    "github.com/stewbawka/go-auth/jwt"
    "github.com/stewbawka/go-auth/controllers"
)

func main() {
    database.Connect()
    defer database.Close()
    database.Migrate()
    jwt.LoadKeypair("/jwt-keypairs")

    r := gin.Default()

    r.GET("/jwks", controllers.GetJwks)

    r.GET("/users", controllers.FindUsers)
    r.GET("/users/:id", controllers.FindUser)
    r.POST("/users", controllers.CreateUser)
    r.PATCH("/users/:id", controllers.UpdateUser)
    r.DELETE("/users/:id", controllers.DeleteUser)

    r.POST("/tokens", controllers.CreateToken)
    r.GET("/tokens/me", controllers.GetTokenByCookie)

    r.Run()
}

