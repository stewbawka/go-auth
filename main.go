package main

import (
    "github.com/gin-gonic/gin"
    "github.com/stewbawka/go-auth/database"
    "github.com/stewbawka/go-auth/event_stream"
    "github.com/stewbawka/go-auth/jwt"
    "github.com/stewbawka/go-auth/controllers"
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
)

func notblank(fl validator.FieldLevel) bool {
    if val, ok := fl.Field().Interface().(*string); ok {
        if val == nil {
            return true
        }
        if *val == "" {
            return false
        }
    }
    return true
}

func main() {
    database.Connect()
    defer database.Close()
    database.Migrate()

    event_stream.Connect()
    defer event_stream.EventStreamConn.Close()

    jwt.LoadKeypair("/jwt-keypairs")

    r := gin.Default()
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
         err := v.RegisterValidation("notblank", notblank)
         if err != nil {
         }
    }
    r.GET("/jwks", controllers.GetJwks)

    r.GET("/users", controllers.FindUsers)
    r.GET("/users/:id", controllers.FindUser)
    r.POST("/users", controllers.CreateUser)
    r.PATCH("/users/:id", controllers.UpdateUser)
    r.DELETE("/users/:id", controllers.DeleteUser)

    r.POST("/tokens", controllers.CreateToken)
    r.GET("/tokens/me", controllers.GetTokenByCookie)
    r.POST("/tokens/invalidate", controllers.InvalidateTokenByCookie)

    r.Run()
}

