package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func ServeJSON(c *fiber.Ctx) error {
	file, err := os.ReadFile("./body.json")
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).SendString(fiber.ErrInternalServerError.Message)
	}
	return c.Send(file)
}
