package main

import (
	"apigateway/config"
	"apigateway/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	cfg, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("cfg", cfg)
		return c.Next()
	})

	routes.Setup(app)
	log.Fatal(app.Listen(":3000"))
}
