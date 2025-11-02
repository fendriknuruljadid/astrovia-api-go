package main

import (
    "github.com/gin-gonic/gin"
    "astrovia-api-go/internal/packages/db"
    "astrovia-api-go/internal/services/user/handlers"
    "astrovia-api-go/internal/middlewares"
)

func main() {
    db.Connect()

    r := gin.Default()
    r.Use(middlewares.RequestID())
    r.Use(middlewares.ErrorHandler())

    r.GET("/users", handlers.GetUsers)
    r.GET("/users/:id", handlers.GetUserByID)
    r.POST("/users", handlers.CreateUser)
    r.PUT("/users/:id", handlers.UpdateUser)
    r.DELETE("/users/:id", handlers.DeleteUser)

    r.Run("127.0.0.1:2001")
}
