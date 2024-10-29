package gitlab

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// ReleasesService handles communication with the releases methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/releases/index.html
type ReleasesService service

// Release represents a project release.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/releases/index.html#list-releases
type Release struct {
	TagName         string    `json:"tag_name"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	DescriptionHTML string    `json:"description_html"`
	CreatedAt       time.Time `json:"created_at"`
	ReleasedAt      time.Time `json:"released_at"`
	Author          struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"author"`
	Commit          Commit `json:"commit"`
	UpcomingRelease bool   `json:"upcoming_release"`
	CommitPath      string `json:"commit_path"`
	TagPath         string `json:"tag_path"`
	Assets          struct {
		Count   int `json:"count"`
		Sources []struct {
			Format string `json:"format"`
			URL    string `json:"url"`
		} `json:"sources"`
		Links []*ReleaseLink `json:"links"`
	} `json:"assets"`
	Links struct {
		ClosedIssueURL     string `json:"closed_issues_url"`
		ClosedMergeRequest string `json:"closed_merge_requests_url"`
		EditURL            string `json:"edit_url"`
		MergedMergeRequest string `json:"merged_merge_requests_url"`
		OpenedIssues       string `json:"opened_issues_url"`
		OpenedMergeRequest string `json:"opened_merge_requests_url"`
		Self               string `json:"self"`
	} `json:"_links"`
}

// ListReleasesOptions represents ListReleases() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/releases/index.html#list-releases
type ListReleasesOptions struct {
	ListOptions `query:",inline"`

	IncludeHTMLDescription *bool `query:"include_html_description,omitempty" json:"include_html_description,omitempty"`
}

// ListReleases gets a pagenated of releases accessible by the authenticated user.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/releases/index.html#list-releases
func (s *ReleasesService) ListReleases(ctx context.Context, projectID string, opts *ListReleasesOptions) ([]*Release, error) {
	apiEndpoint := fmt.Sprintf("projects/%s/releases", projectID)
	var v []*Release
	if _, err := s.client.InvokeByCredential(ctx, http.MethodGet, apiEndpoint, opts, &v); err != nil {
		return nil, err
	}
	return v, nil
}
