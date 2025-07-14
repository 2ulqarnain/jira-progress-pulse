package main

import (
	"jira-progress-pulse/internal/config"
	"jira-progress-pulse/internal/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.Load()
	app := fiber.New()
	// allowing all origins for now - TODO: allow only necessary origins
	app.Use(cors.New())
	app.Use(logger.New())

	issueGroup := app.Group("/issues")
	issueGroup.Get("/:issueID", controllers.GetIssueByID)

	app.Static("/", "./res/home.html")
	app.Listen(":8001")

}
