package repository

import (
	"encoding/json"
	"fmt"
	"jira-progress-pulse/internal/config"
	"jira-progress-pulse/internal/dto"
	"net/http"
)

type JiraRepository struct {
	BaseURL  string // e.g., "https://your-company.atlassian.net"
	Username string // Your Jira email (e.g., "user@company.com")
	APIToken string // Generated in Jira profile settings
}

func NewJiraRepository(baseURL, username, apiToken string) *JiraRepository {
	return &JiraRepository{
		BaseURL:  baseURL,
		Username: username,
		APIToken: apiToken,
	}
}

// Fetch a single issue by ID/key (e.g., "DTS-1343")
func GetIssue(issueID string) (*dto.JiraIssueResponse, error) {
	url := fmt.Sprintf("%s/rest/api/3/issue/%s", config.Cfg.Jira.BaseURL, issueID)

	// 1. Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// 2. Set Basic Auth headers
	req.SetBasicAuth(config.Cfg.Jira.Username, config.Cfg.Jira.APIToken)
	req.Header.Set("Accept", "application/json")

	// 3. Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Jira API request failed: %v", err)
	}
	defer resp.Body.Close()

	// 4. Handle non-200 responses
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Jira API error: %s", resp.Status)
	}

	// 5. Parse response
	var issue dto.JiraIssueResponse
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, fmt.Errorf("failed to decode Jira response: %v", err)
	}

	return &issue, nil
}
