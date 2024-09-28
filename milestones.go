package gitlab

import "time"

// Milestone represents a GitLab milestone.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/milestones.html
type Milestone struct {
	ID          int        `json:"id"`
	IID         int        `json:"iid"`
	GroupID     int        `json:"group_id"`
	ProjectID   int        `json:"project_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StartDate   *time.Time `json:"start_date"`
	DueDate     *time.Time `json:"due_date"`
	State       string     `json:"state"`
	WebURL      string     `json:"web_url"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at"`
	Expired     bool       `json:"expired"`
}
