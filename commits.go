package gitlab

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// CommitsService
// GitLab API docs: https://docs.gitlab.com/ee/api/commits.html
type CommitsService service

type Commit struct {
	ID             string            `json:"id"`
	ShortID        string            `json:"short_id"`
	Title          string            `json:"title"`
	AuthorName     string            `json:"author_name"`
	AuthorEmail    string            `json:"author_email"`
	AuthoredDate   time.Time         `json:"authored_date"`
	CommitterName  string            `json:"committer_name"`
	CommitterEmail string            `json:"committer_email"`
	CommittedDate  time.Time         `json:"committed_date"`
	CreatedAt      time.Time         `json:"created_at"`
	Message        string            `json:"message"`
	ParentIDs      []string          `json:"parent_ids"`
	Stats          *CommitStats      `json:"stats"`
	Status         *BuildStateValue  `json:"status"`
	LastPipeline   *PipelineInfo     `json:"last_pipeline"`
	ProjectID      int               `json:"project_id"`
	Trailers       map[string]string `json:"trailers"`
	WebURL         string            `json:"web_url"`
}

type CommitStats struct {
	Additions int64 `json:"additions"`
	Deletions int64 `json:"deletions"`
	Total     int64 `json:"total"`
}

// ListCommitsOptions represents the available ListCommits() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/commits.html#list-repository-commits
type ListCommitsOptions struct {
	ListOptions `query:",inline"`

	RefName     *string    `query:"ref_name,omitempty"`
	Since       *time.Time `query:"since,omitempty"`
	Until       *time.Time `query:"until,omitempty"`
	Path        *string    `query:"path,omitempty"`
	Author      *string    `query:"author,omitempty"`
	All         *bool      `query:"all,omitempty"`
	WithStats   *bool      `query:"with_stats,omitempty"`
	FirstParent *bool      `query:"first_parent,omitempty"`
	Trailers    *bool      `query:"trailers,omitempty"`
}

// ListCommits gets a list of repository commits in a project.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/commits.html#list-repository-commits
func (s *CommitsService) ListCommits(ctx context.Context, projectId string, opts *ListCommitsOptions) ([]*Commit, error) {
	apiEndpoint := fmt.Sprintf("projects/%s/repository/commits", projectId)
	var v []*Commit
	if _, err := s.client.InvokeByCredential(ctx, http.MethodGet, apiEndpoint, opts, &v); err != nil {
		return nil, err
	}
	return v, nil
}
