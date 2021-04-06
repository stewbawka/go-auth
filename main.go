package main

import (
    "github.com/gin-gonic/gin"
    "github.com/stewbawka/go-auth/database"
    "github.com/stewbawka/go-auth/controllers"
)

func main() {
    database.Connect()
    database.Migrate()
    r := gin.Default()

    r.GET("/users", controllers.FindUsers)
    r.GET("/users/:id", controllers.FindUser)
    r.POST("/users", controllers.CreateUser)
    r.PATCH("/users/:id", controllers.UpdateUser)
    r.DELETE("/users/:id", controllers.DeleteUser)

    r.Run()
}

