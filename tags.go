package gitlab

import (
	"context"
	"fmt"
	"net/http"
)

// TagsService
// GitLab API docs: https://docs.gitlab.com/ee/api/tags.html
type TagsService service

type Tag struct {
	Commit    *Commit      `json:"commit"`
	Release   *ReleaseNote `json:"release"`
	Name      string       `json:"name"`
	Message   string       `json:"message"`
	Protected bool         `json:"protected"`
	Target    string       `json:"target"`
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
	*ListOptions `query:",inline"`
	OrderBy      *string `query:"order_by,omitempty"`
	Search       *string `query:"search,omitempty"`
}

func (s *TagsService) ListTags(ctx context.Context, projectId string, opts *ListTagsOptions) ([]*Tag, error) {
	apiEndpoint := fmt.Sprintf("/api/v4/projects/%s/repository/tags", projectId)
	var v []*Tag
	if err := s.client.InvokeByCredential(ctx, http.MethodGet, apiEndpoint, opts, &v); err != nil {
		return nil, err
	}
	return v, nil
}
