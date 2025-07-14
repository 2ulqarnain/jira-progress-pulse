package dto

// DTO for Jira get issue by id api response
type JiraIssueResponse struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Fields struct {
		StatusCategoryChangeDate string `json:"statuscategorychangedate"`
		Status                   struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"status"`
		Summary  string `json:"summary"`
		Assignee struct {
			DisplayName string            `json:"displayName"`
			AvatarUrls  map[string]string `json:"avatarUrls"`
		} `json:"assignee"`
		TimeTracking struct {
			OriginalEstimate string `json:"originalEstimate"`
		} `json:"timetracking"`
	} `json:"fields"`
}

type GetIssueResponse struct {
	JiraID string `json:"jiraId"`
	ID     string `json:"id"`
	Status struct {
		StatusName        string `json:"name"`
		StatusDescription string `json:"description"`
		StatusChangedDate string `json:"changedOn"`
	} `json:"status"`
	LastCommitTime string `json:"lastCommitTime"`
	Assignee       struct {
		DisplayName string `json:"name"`
		// TODO: implement logic that sends only the required avatar url
		AvatarUrls map[string]string `json:"avatarUrls"`
	} `json:"assignee"`
	OriginalEstimate string `json:"originalEstimate"`
	Summary          string `json:"summary"`
	PRStatus         string `json:"prStatus"`
	TimeInProgress   string `json:"timeInProgress"` // e.g., "3h45m"
}
