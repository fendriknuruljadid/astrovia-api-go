package main

import (
	routesV1 "app/cmd/services/v1/astro-zenith/routes"
	"app/internal/middlewares"
	"app/internal/packages/db"
	"app/internal/packages/redis"

	"github.com/gin-gonic/gin"
)

func main() {

	db.Connect()
	redis.InitRedis()
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.SetTrustedProxies([]string{
		"127.0.0.1",  // kalau Fiber & Gin satu host
		"10.0.0.0/8", // docker / swarm network
		"172.16.0.0/12",
		"192.168.0.0/16",
	})

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
