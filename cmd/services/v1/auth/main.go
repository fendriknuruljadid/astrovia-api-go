package main

import (
    "github.com/gin-gonic/gin"
    "app/internal/packages/db"
    auth "app/internal/services/v1/auth/handlers"
    "app/internal/middlewares"
)

func main() {
    db.Connect()

    r := gin.Default()
    r.Use(middlewares.RequestID())
    r.Use(middlewares.ErrorHandler())

    r.POST("/v1/generate-token", auth.Auth)

    r.Run("127.0.0.1:2002")
}
