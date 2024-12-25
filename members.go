package gitlab

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// MembersService
// GitLab API docs: https://docs.gitlab.com/ee/api/members.html
type MembersService service

// Member represents a GitLab member.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/members.html
type Member struct {
	ID                int                      `json:"id"`
	Username          string                   `json:"username"`
	Name              string                   `json:"name"`
	State             string                   `json:"state"`
	AvatarURL         string                   `json:"avatar_url"`
	WebURL            string                   `json:"web_url"`
	CreatedAt         *time.Time               `json:"created_at"`
	ExpiresAt         *Date                    `json:"expires_at"`
	AccessLevel       AccessLevelValue         `json:"access_level"`
	Email             string                   `json:"email,omitempty"`
	GroupSAMLIdentity *GroupMemberSAMLIdentity `json:"group_saml_identity"`
	MemberRole        *MemberRole              `json:"member_role"`
}

// GroupMemberSAMLIdentity represents the SAML Identity link for the group member.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/members.html#list-all-members-of-a-group-or-project
type GroupMemberSAMLIdentity struct {
	ExternUID      string `json:"extern_uid"`
	Provider       string `json:"provider"`
	SAMLProviderID int    `json:"saml_provider_id"`
}

// ListMembersOptions
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/members.html#list-all-members-of-a-group-or-project
type ListMembersOptions struct {
	ListOptions
	Query        *string `url:"query,omitempty"`
	UserIDs      *[]int  `url:"user_ids[],omitempty"`
	SkipUsers    *[]int  `url:"skip_users[],omitempty"`
	ShowSeatInfo bool    `url:"show_seat_info,omitempty"`
}

// ListGroupMembers get a list of group members viewable by the authenticated
// user. Inherited members through ancestor groups are not included.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/members.html#list-all-members-of-a-group-or-project
func (s *MembersService) ListGroupMembers(ctx context.Context, gid string, opts *ListMembersOptions) ([]*Member, *Page, error) {
	var reply []*Member
	u := fmt.Sprintf("groups/%s/members", gid)
	resp, err := s.client.InvokeWithCredential(ctx, http.MethodGet, u, opts, &reply)
	if err != nil {
		return nil, nil, err
	}

	return reply, NewPage(opts, resp), nil
}

// ListAllGroupMembers get a list of group members viewable by the authenticated
// user. Returns a list including inherited members through ancestor groups.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/members.html#list-all-members-of-a-group-or-project-including-inherited-and-invited-members
func (s *MembersService) ListAllGroupMembers(ctx context.Context, gid string, opts *ListMembersOptions) ([]*Member, *Page, error) {
	var reply []*Member
	u := fmt.Sprintf("groups/%s/members/all", gid)
	resp, err := s.client.InvokeWithCredential(ctx, http.MethodGet, u, opts, &reply)
	if err != nil {
		return nil, nil, err
	}

	return reply, NewPage(opts, resp), nil
}

// ListProjectMembers gets a list of a project's team members viewable by the
// authenticated user. Returns only direct members and not inherited members
// through ancestors groups.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/members.html#list-all-members-of-a-group-or-project
func (s *MembersService) ListProjectMembers(ctx context.Context, gid string, opts *ListMembersOptions) ([]*Member, *Page, error) {
	var reply []*Member
	u := fmt.Sprintf("projects/%s/members", gid)
	resp, err := s.client.InvokeWithCredential(ctx, http.MethodGet, u, opts, &reply)
	if err != nil {
		return nil, nil, err
	}

	return reply, NewPage(opts, resp), nil
}

// ListAllProjectMembers gets a list of a project's team members viewable by the
// authenticated user. Returns a list including inherited members through
// ancestor groups.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/members.html#list-all-members-of-a-group-or-project-including-inherited-and-invited-members
func (s *MembersService) ListAllProjectMembers(ctx context.Context, gid string, opts *ListMembersOptions) ([]*Member, *Page, error) {
	var reply []*Member
	u := fmt.Sprintf("projects/%s/members/all", gid)
	resp, err := s.client.InvokeWithCredential(ctx, http.MethodGet, u, opts, &reply)
	if err != nil {
		return nil, nil, err
	}

	return reply, NewPage(opts, resp), nil
}
