package controllers

import (
	"jira-progress-pulse/internal/dto"
	"jira-progress-pulse/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func GetIssueByID(c *fiber.Ctx) error {
	issueIdParam := c.Params("issueID")
	response, err := repository.GetIssue(issueIdParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": err.Error()})
	}
	c.JSON(dto.GetIssueResponse{
		JiraID: response.ID,
		ID:     response.Key,
		Status: struct {
			StatusName        string `json:"name"`
			StatusDescription string `json:"description"`
			StatusChangedDate string `json:"changedOn"`
		}{
			StatusName:        response.Fields.Status.Name,
			StatusDescription: response.Fields.Status.Description,
			StatusChangedDate: response.Fields.StatusCategoryChangeDate,
		},
		Assignee: struct {
			DisplayName string            `json:"name"`
			AvatarUrls  map[string]string `json:"avatarUrls"`
		}{
			DisplayName: response.Fields.Assignee.DisplayName,
			AvatarUrls:  response.Fields.Assignee.AvatarUrls,
		},
		OriginalEstimate: response.Fields.TimeTracking.OriginalEstimate,
		Summary:          response.Fields.Summary,
	})
	return nil
}
