package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/swagger"
	_ "astrovia-api-go/cmd/gateway/docs"
	_ "astrovia-api-go/cmd/gateway/routes"
	"astrovia-api-go/internal/middlewares"
)


// =================== Main Fiber App ===================

// @title Astrovia Gateway API
// @version 1.1
// @description API Gateway Astrovia AI.
// @BasePath /

// @securityDefinitions.apikey X-Signature
// @in header
// @name X-Signature
// @description Signature HMAC SHA256 untuk validasi integritas request

// @securityDefinitions.apikey X-Timestamp
// @in header
// @name X-Timestamp
// @description Timestamp UTC format RFC3339 (contoh: 2025-11-03T09:10:00Z)

// @contact.name API Support
// @contact.email dev@astrovia.id


func main() {
	app := fiber.New()

	// CORS supaya Swagger UI bisa fetch
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	// Swagger UI
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use("/users", middlewares.SignatureClientMiddleware())
	// Proxy semua request /users ke service Gin
	app.All("/users/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, "http://localhost:2001"+c.OriginalURL())
	})

	// Jalankan server gateway
	app.Listen(":2000")
}
