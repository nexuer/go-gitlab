package gitlab

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// TagsService
// GitLab API docs: https://docs.gitlab.com/ee/api/tags.html
type TagsService service

type Tag struct {
	Commit    Commit       `json:"commit"`
	Release   *ReleaseNote `json:"release"`
	Name      string       `json:"name"`
	Message   string       `json:"message"`
	Protected bool         `json:"protected"`
	Target    string       `json:"target"`
	CreatedAt time.Time    `json:"created_at"`
}

type ReleaseNote struct {
	TagName     string `json:"tag_name"`
	Description string `json:"description"`
}

// ListTagsOptions represents the available ListTags() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/tags.html#list-project-repository-tags
type ListTagsOptions struct {
	ListOptions `query:",inline"`

	Search *string `query:"search,omitempty"`
}

func (s *TagsService) ListTags(ctx context.Context, projectId string, opts *ListTagsOptions) ([]*Tag, *PageInfo, error) {
	apiEndpoint := fmt.Sprintf("projects/%s/repository/tags", projectId)
	var v []*Tag
	resp, err := s.client.InvokeWithCredential(ctx, http.MethodGet, apiEndpoint, opts, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, opts.ListOptions.ParsePageInfo(resp), nil
}
