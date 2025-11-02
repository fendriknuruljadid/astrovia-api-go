package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/swagger"
	_ "astrovia-api-go/cmd/gateway/docs"
	_ "astrovia-api-go/cmd/gateway/routes"
)


// =================== Main Fiber App ===================
func main() {
	app := fiber.New()

	// CORS supaya Swagger UI bisa fetch
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	// Swagger UI
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Proxy semua request /users ke service Gin
	app.All("/users/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, "http://localhost:2001"+c.OriginalURL())
	})

	// Jalankan server gateway
	app.Listen(":2000")
}
