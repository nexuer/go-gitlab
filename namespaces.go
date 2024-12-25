package gitlab

import (
	"context"
	"net/http"
)

// NamespacesService handles communication with the namespace related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/namespaces.html
type NamespacesService service

// Namespace represents a GitLab namespace.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/namespaces.html
type Namespace struct {
	ID                          int     `json:"id"`
	Name                        string  `json:"name"`
	Path                        string  `json:"path"`
	Kind                        string  `json:"kind"`
	FullPath                    string  `json:"full_path"`
	ParentID                    int     `json:"parent_id"`
	AvatarURL                   *string `json:"avatar_url"`
	WebURL                      string  `json:"web_url"`
	MembersCountWithDescendants int     `json:"members_count_with_descendants"`
	BillableMembersCount        int     `json:"billable_members_count"`
	Plan                        string  `json:"plan"`
	TrialEndsOn                 *Date   `json:"trial_ends_on"`
	Trial                       bool    `json:"trial"`
	MaxSeatsUsed                *int    `json:"max_seats_used"`
	SeatsInUse                  *int    `json:"seats_in_use"`
}

// ListNamespacesOptions represents the available ListNamespaces() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/namespaces.html#list-namespaces
type ListNamespacesOptions struct {
	ListOptions `query:",inline,omitempty"`

	Search       *string `query:"search,omitempty"`
	OwnedOnly    *bool   `query:"owned_only,omitempty"`
	TopLevelOnly *bool   `query:"top_level_only,omitempty"`
}

// ListNamespaces gets a list of projects accessible by the authenticated user.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/namespaces.html#list-namespaces
func (s *NamespacesService) ListNamespaces(ctx context.Context, opts *ListNamespacesOptions) ([]*Namespace, *Page, error) {
	var reply []*Namespace
	resp, err := s.client.InvokeWithCredential(ctx, http.MethodGet, "namespaces", opts, &reply)
	if err != nil {
		return nil, nil, err
	}

	return reply, NewPage(opts, resp), nil
}
