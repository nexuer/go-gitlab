package gitlab

import (
	"context"
	"fmt"
	"net/http"
)

// BranchesService
// GitLab API docs: https://docs.gitlab.com/ee/api/branches.html
type BranchesService service

type Branch struct {
	Commit             Commit `json:"commit"`
	Name               string `json:"name"`
	Protected          bool   `json:"protected"`
	Merged             bool   `json:"merged"`
	Default            bool   `json:"default"`
	CanPush            bool   `json:"can_push"`
	DevelopersCanPush  bool   `json:"developers_can_push"`
	DevelopersCanMerge bool   `json:"developers_can_merge"`
	WebURL             string `json:"web_url"`
}

// ListBranchesOptions represents the available ListBranches() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/branches.html#list-repository-branches
type ListBranchesOptions struct {
	ListOptions `query:",inline"`
	Search      *string `query:"search,omitempty"`
	Regex       *string `query:"regex,omitempty"`
}

// ListBranches gets a list of repository branches from a project, sorted by
// name alphabetically.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/branches.html#list-repository-branches
func (s *BranchesService) ListBranches(ctx context.Context, projectID string, opts *ListBranchesOptions) ([]*Branch, error) {
	apiEndpoint := fmt.Sprintf("projects/%s/repository/branches", projectID)
	var v []*Branch
	if _, err := s.client.InvokeByCredential(ctx, http.MethodGet, apiEndpoint, opts, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// CreateBranchOptions represents the available CreateBranch() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/branches.html#create-repository-branch
type CreateBranchOptions struct {
	Branch *string `json:"branch,omitempty"`
	Ref    *string `json:"ref,omitempty"`
}

func (s *BranchesService) CreateBranch(ctx context.Context, projectId string, opts *CreateBranchOptions) (*Branch, error) {
	apiEndpoint := fmt.Sprintf("projects/%s/repository/branches", projectId)
	var v *Branch
	if _, err := s.client.InvokeByCredential(ctx, http.MethodPost, apiEndpoint, opts, &v); err != nil {
		return nil, err
	}
	return v, nil
}

func (s *BranchesService) DeleteBranch(ctx context.Context, projectId, branch string) error {
	apiEndpoint := fmt.Sprintf("projects/%s/repository/branches/%s", projectId, branch)
	if _, err := s.client.InvokeByCredential(ctx, http.MethodDelete, apiEndpoint, nil, nil); err != nil {
		return err
	}
	return nil
}

func (s *BranchesService) DeleteMergedBranches(ctx context.Context, projectId string) error {
	apiEndpoint := fmt.Sprintf("projects/%s/repository/merged_branches", projectId)
	if _, err := s.client.InvokeByCredential(ctx, http.MethodDelete, apiEndpoint, nil, nil); err != nil {
		return err
	}
	return nil
}
