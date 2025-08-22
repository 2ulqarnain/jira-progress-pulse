package main

import (
	"example.com/m.v2/config"
	"example.com/m.v2/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	config.Load()

	app := fiber.New()

	app.Use(cors.New())

	app.Post("/jira-issues-transition-webhook", controllers.JiraIssueTransitionWebhook)
	app.Get("/body-json", controllers.ServeJSON)

	app.Listen(":8000")
}
