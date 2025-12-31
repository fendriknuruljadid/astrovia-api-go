package main

import (
	_ "app/cmd/gateway/docs"
	_ "app/cmd/gateway/swagger/v1"
	"app/internal/middlewares"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/swagger"
)

// =================== Main Fiber App ===================

// @title Astrovia Gateway API
// @version 1.1
// @description API Gateway Astrovia AI.
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Format: Bearer <token>

// @securityDefinitions.apikey X-Signature
// @in header
// @name X-Signature
// @description Signature HMAC SHA256 untuk validasi integritas request

// @securityDefinitions.apikey X-Timestamp
// @in header
// @name X-Timestamp
// @description Timestamp UTC format RFC3339 (contoh: 2025-11-03T09:10:00Z)

// @securityDefinitions.apikey X-DeviceId
// @in header
// @name X-DeviceId
// @description DeviceId untuk validasi request

// @contact.name API Support
// @contact.email dev@astrovia.id
var (
	userServiceURL   = os.Getenv("USER_SERVICE_URL")
	globalServiceURL = os.Getenv("GLOBAL_SERVICE_URL")
	astroZenithURL   = os.Getenv("ASTRO_ZENITH_URL")
)

func main() {
	app := fiber.New()

	// CORS supaya Swagger UI bisa fetch
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	// Swagger UI
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Semua route users v1
	apiV1 := app.Group("/v1")

	// Auth & User Service
	userGrp := apiV1.Group("/users")
	userGrp.Use(middlewares.SignatureClientMiddleware())
	userGrp.All("/*", func(c *fiber.Ctx) error {
		target := userServiceURL + c.OriginalURL()
		return proxy.Do(c, target)
	})

	authGrp := apiV1.Group("/auth")
	authGrp.Use(middlewares.SignatureClientMiddleware())
	authGrp.All("/*", func(c *fiber.Ctx) error {
		target := userServiceURL + c.OriginalURL()
		return proxy.Do(c, target)
	})

	// global service
	agents := apiV1.Group("/agents")
	agents.Use(middlewares.SignatureClientMiddleware())
	agents.All("/*", func(c *fiber.Ctx) error {
		target := globalServiceURL + c.OriginalURL()
		return proxy.Do(c, target)
	})

	pricing := apiV1.Group("/pricing")
	pricing.Use(middlewares.SignatureClientMiddleware())
	pricing.All("/*", func(c *fiber.Ctx) error {
		target := globalServiceURL + c.OriginalURL()
		return proxy.Do(c, target)
	})

	userAgents := apiV1.Group("/user-agents")
	userAgents.Use(middlewares.SignatureClientMiddleware())
	userAgents.All("/*", func(c *fiber.Ctx) error {
		target := globalServiceURL + c.OriginalURL()
		return proxy.Do(c, target)
	})

	// Astro zenith service
	astroZenithGrp := apiV1.Group("/astro-zenith")
	astroZenithGrp.Use(middlewares.SignatureClientMiddleware())

	autoClipGrp := astroZenithGrp.Group("/auto-clip")
	autoClipGrp.All("/*", func(c *fiber.Ctx) error {
		target := astroZenithURL + c.OriginalURL()
		return proxy.Do(c, target)
	})

	autoCaption := astroZenithGrp.Group("/auto-caption")
	autoCaption.All("/*", func(c *fiber.Ctx) error {
		target := astroZenithURL + c.OriginalURL()
		return proxy.Do(c, target)
	})
	// Jalankan server gateway
	app.Listen(":2000")
}
