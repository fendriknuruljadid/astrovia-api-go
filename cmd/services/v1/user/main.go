package main

import (
    "github.com/gin-gonic/gin"
    "astrovia-api-go/internal/packages/db"
    "astrovia-api-go/internal/middlewares"
    "astrovia-api-go/cmd/services/v1/user/routes"
)

func main() {

    db.Connect()

    r := gin.Default()

    // Global middleware
    r.Use(middlewares.RequestID())
    r.Use(middlewares.ErrorHandler())

    // ================================================
    // PUBLIC group (no JWT)
    // ================================================
    publicV1 := r.Group("/v1")
    routesV1.RegisterPublicRoutes(publicV1)


    protectedV1 := r.Group("/v1")
    protectedV1.Use(middlewares.JWTAuth())   
    routesV1.RegisterProtectedRoutes(protectedV1)


    r.Run(":2001")
}
