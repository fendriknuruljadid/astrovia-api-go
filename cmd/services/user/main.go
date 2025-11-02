package main

import (
    "github.com/gin-gonic/gin"
    "astrovia-api-go/internal/package/db"
    "astrovia-api-go/internal/services/user/handlers"
)

func main() {
    db.Connect()

    r := gin.Default()

    r.GET("/users", handlers.GetUsers)
    r.GET("/users/:id", handlers.GetUserByID)
    r.POST("/users", handlers.CreateUser)
    r.PUT("/users/:id", handlers.UpdateUser)
    r.DELETE("/users/:id", handlers.DeleteUser)

    r.Run("127.0.0.1:2001")
}
