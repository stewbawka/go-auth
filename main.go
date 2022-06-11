package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/stewbawka/go-auth/database"
    "github.com/stewbawka/go-auth/jwt"
    "github.com/stewbawka/go-auth/controllers"
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
)

func notblank(fl validator.FieldLevel) bool {
    fmt.Printf("validator")
    if val, ok := fl.Field().Interface().(*string); ok {
        fmt.Printf("inside validator")
        fmt.Printf("<Value:%d>\n", val)
        if val == nil {
            fmt.Printf("there")
            return true
        }
        if *val == "" {
            fmt.Printf("here")
            return false
        }
    } else {
        fmt.Printf("unable to cast")
    }
    return true
}

func main() {
    database.Connect()
    defer database.Close()
    database.Migrate()
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

