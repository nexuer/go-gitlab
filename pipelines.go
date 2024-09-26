package gitlab

import "time"

// Pipeline represents a GitLab pipeline.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/pipelines.html
type Pipeline struct {
	ID             int             `json:"id"`
	IID            int             `json:"iid"`
	ProjectID      int             `json:"project_id"`
	Status         string          `json:"status"`
	Source         string          `json:"source"`
	Ref            string          `json:"ref"`
	SHA            string          `json:"sha"`
	BeforeSHA      string          `json:"before_sha"`
	Tag            bool            `json:"tag"`
	YamlErrors     string          `json:"yaml_errors"`
	User           *BasicUser      `json:"user"`
	UpdatedAt      *time.Time      `json:"updated_at"`
	CreatedAt      *time.Time      `json:"created_at"`
	StartedAt      *time.Time      `json:"started_at"`
	FinishedAt     *time.Time      `json:"finished_at"`
	CommittedAt    *time.Time      `json:"committed_at"`
	Duration       int             `json:"duration"`
	QueuedDuration int             `json:"queued_duration"`
	Coverage       string          `json:"coverage"`
	WebURL         string          `json:"web_url"`
	DetailedStatus *DetailedStatus `json:"detailed_status"`
}

// DetailedStatus contains detailed information about the status of a pipeline.
type DetailedStatus struct {
	Icon         string `json:"icon"`
	Text         string `json:"text"`
	Label        string `json:"label"`
	Group        string `json:"group"`
	Tooltip      string `json:"tooltip"`
	HasDetails   bool   `json:"has_details"`
	DetailsPath  string `json:"details_path"`
	Illustration struct {
		Image string `json:"image"`
	} `json:"illustration"`
	Favicon string `json:"favicon"`
}

// PipelineInfo shows the basic entities of a pipeline, mostly used as fields
// on other assets, like Commit.
type PipelineInfo struct {
	ID        int        `json:"id"`
	IID       int        `json:"iid"`
	ProjectID int        `json:"project_id"`
	Status    string     `json:"status"`
	Source    string     `json:"source"`
	Ref       string     `json:"ref"`
	SHA       string     `json:"sha"`
	WebURL    string     `json:"web_url"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
}
