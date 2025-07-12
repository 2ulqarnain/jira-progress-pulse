package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	// allowing all origins for now - TODO: allow only necessary origins
	app.Use(cors.New())
	app.Use(logger.New())

	app.Static("/", "./res/home.html")
	app.Listen(":8001")

}
