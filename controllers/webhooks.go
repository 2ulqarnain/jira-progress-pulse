package controllers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func JiraIssueTransitionWebhook(c *fiber.Ctx) error {
	bodyInBytes := c.Body()

	if len(bodyInBytes) == 0 {
		fmt.Println("Body is empty")
		return c.SendStatus(fiber.StatusOK)
	}

	err := os.WriteFile("./body.json", bodyInBytes, 0644)
	if err != nil {
		fmt.Println("Couldn't write to file")
	}
	fmt.Println("File wrote successfully !")
	return c.SendStatus(fiber.StatusOK)
}
