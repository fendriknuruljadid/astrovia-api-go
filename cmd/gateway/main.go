package main

import (
	_ "app/cmd/gateway/docs"
	_ "app/cmd/gateway/swagger/v1"
	"app/internal/middlewares"

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

	apiV1 := app.Group("/v1")

	// videoGrp.Post("/video/process", func(c *fiber.Ctx) error {
	// 	var data Job
	// 	if err := c.BodyParser(&data); err != nil {
	// 		return err
	// 	}

	// 	jobData, _ := json.Marshal(data)
	// 	redisClient.LPush(context.Background(), "video_jobs", jobData)

	// 	return c.JSON(fiber.Map{"status": "queued", "job_id": data.JobID})
	// })

	// Semua route users v1
	userGrp := apiV1.Group("/users")
	userGrp.Use(middlewares.SignatureClientMiddleware())
	userGrp.All("/*", func(c *fiber.Ctx) error {
		target := "http://localhost:2001" + c.OriginalURL()
		return proxy.Do(c, target)
	})

	authGrp := apiV1.Group("/generate-token")
	authGrp.Use(middlewares.SignatureClientMiddleware())
	authGrp.All("/*", func(c *fiber.Ctx) error {
		target := "http://localhost:2002" + c.OriginalURL()
		return proxy.Do(c, target)
	})

	videoGrp := apiV1.Group("/video")
	videoGrp.Use(middlewares.SignatureClientMiddleware())
	videoGrp.All("/*", func(c *fiber.Ctx) error {
		target := "http://localhost:2003" + c.OriginalURL()
		return proxy.Do(c, target)
	})

	youtuberGrp := apiV1.Group("/youtuber")
	youtuberGrp.Use(middlewares.SignatureClientMiddleware())
	autoShortCutGrp := youtuberGrp.Group("/auto-short")
	autoShortCutGrp.All("/*", func(c *fiber.Ctx) error {
		target := "http://localhost:2005" + c.OriginalURL()
		return proxy.Do(c, target)
	})

	// Jalankan server gateway
	app.Listen(":2000")
}
