package gitlab

// TimeStats represents the time estimates and time spent for an issue.
//
// GitLab docs: https://docs.gitlab.com/ee/workflow/time_tracking.html
type TimeStats struct {
	HumanTimeEstimate   string `json:"human_time_estimate"`
	HumanTotalTimeSpent string `json:"human_total_time_spent"`
	TimeEstimate        int    `json:"time_estimate"`
	TotalTimeSpent      int    `json:"total_time_spent"`
}
