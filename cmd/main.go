package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/hazitgi/graphql-with-fiber/apis/routes"
	"github.com/hazitgi/graphql-with-fiber/config"
	"github.com/joho/godotenv"
)

var httpPort *string

func init() {
	httpPort = flag.String("http-port", "8000", "http port")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading.env file")
	}
}

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader:  "LoomERP API",
		AppName:       "LoomERP v1.0.0",
		CaseSensitive: true,
	})
	app.Use(healthcheck.New())
	app.Use(logger.New())
	config.ConnectDB()

	app.Get("/metrics", monitor.New(monitor.Config{Title: "LoomERP Metrics Page"}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		result := map[string]interface{}{
			"message": "Hello World",
			"version": "v3.0.0",
			"status":  200,
			"data": map[string]interface{}{
				"first_name": "hazi",
				"last_name":  "tgi",
				"email":      "hazi.tgi@gmail.com",
				"age":        25,
				"address": map[string]interface{}{
					"city":    "Baku",
					"country": "Azerbaijan",
				},
			},
		}

		ctx.Set("Content-Type", "application/graphql-response+json")
		return ctx.JSON(result)
	})

	routes.NewUserRoutes().RegisterUserRoutes(app)

	if err := app.Listen(":" + *httpPort); err != nil {
		log.Fatalf("Error in starting server : %v", err)
	}
}
