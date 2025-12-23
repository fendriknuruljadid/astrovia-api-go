package main

import (
	routesV1 "app/cmd/services/v1/astro-zenith/routes"
	"app/internal/middlewares"
	"app/internal/packages/db"

	"github.com/gin-gonic/gin"
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
	publicV1 := r.Group("/v1/astro-zenith")
	routesV1.RegisterPublicRoutes(publicV1)

	protectedV1 := r.Group("/v1/astro-zenith")
	protectedV1.Use(middlewares.JWTAuth())
	routesV1.RegisterProtectedRoutes(protectedV1)

	r.Run(":2003")
}
