package gitlab

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// MilestonesService handles communication with the milestone related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/milestones.html
type MilestonesService service

// Milestone represents a GitLab milestone.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/milestones.html
type Milestone struct {
	ID          int       `json:"id"`
	IID         int       `json:"iid"`
	GroupID     int       `json:"group_id"`
	ProjectID   int       `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   *Date     `json:"start_date"`
	DueDate     *Date     `json:"due_date"`
	State       string    `json:"state"`
	WebURL      string    `json:"web_url"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	Expired     bool      `json:"expired"`
}

// ListMilestonesOptions represents the available ListMilestones() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/milestones.html#list-project-milestones
type ListMilestonesOptions struct {
	ListOptions `query:",inline"`

	IIDs                    *[]int          `query:"iids[],omitempty" json:"iids,omitempty"`
	Title                   *string         `query:"title,omitempty" json:"title,omitempty"`
	State                   *MilestoneState `query:"state,omitempty" json:"state,omitempty"`
	Search                  *string         `query:"search,omitempty" json:"search,omitempty"`
	IncludeParentMilestones *bool           `query:"include_parent_milestones,omitempty" json:"include_parent_milestones,omitempty"`
	IncludeAncestors        *bool           `query:"include_ancestors,omitempty" json:"include_ancestors,omitempty"`
	UpdatedBefore           *time.Time      `query:"updated_before,omitempty" json:"updated_before,omitempty"`
	UpdatedAfter            *time.Time      `query:"updated_after,omitempty" json:"updated_after,omitempty"`
}

// ListMilestones returns a list of project milestones.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/milestones.html#list-project-milestones
func (s *MilestonesService) ListMilestones(ctx context.Context, pid string, opts *ListMilestonesOptions) ([]*Milestone, error) {
	u := fmt.Sprintf("projects/%s/milestones", pid)

	var milestones []*Milestone
	if _, err := s.client.InvokeByCredential(ctx, http.MethodGet, u, opts, &milestones); err != nil {
		return nil, err
	}
	return milestones, nil

}
