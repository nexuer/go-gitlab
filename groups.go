package gitlab

import (
	"context"
	"net/http"
	"time"
)

// GroupsService handles communication with the group related methods of
// the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/groups.html
type GroupsService service

// Group represents a GitLab group.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/groups.html
type Group struct {
	ID                              int                        `json:"id"`
	Name                            string                     `json:"name"`
	Path                            string                     `json:"path"`
	Description                     string                     `json:"description"`
	MembershipLock                  bool                       `json:"membership_lock"`
	Visibility                      VisibilityValue            `json:"visibility"`
	LFSEnabled                      bool                       `json:"lfs_enabled"`
	DefaultBranch                   string                     `json:"default_branch"`
	DefaultBranchProtectionDefaults *BranchProtectionDefaults  `json:"default_branch_protection_defaults"`
	AvatarURL                       string                     `json:"avatar_url"`
	WebURL                          string                     `json:"web_url"`
	RequestAccessEnabled            bool                       `json:"request_access_enabled"`
	RepositoryStorage               string                     `json:"repository_storage"`
	FullName                        string                     `json:"full_name"`
	FullPath                        string                     `json:"full_path"`
	FileTemplateProjectID           int                        `json:"file_template_project_id"`
	ParentID                        int                        `json:"parent_id"`
	Projects                        []*Project                 `json:"projects"`
	Statistics                      *Statistics                `json:"statistics"`
	CustomAttributes                []*CustomAttribute         `json:"custom_attributes"`
	ShareWithGroupLock              bool                       `json:"share_with_group_lock"`
	RequireTwoFactorAuth            bool                       `json:"require_two_factor_authentication"`
	TwoFactorGracePeriod            int                        `json:"two_factor_grace_period"`
	ProjectCreationLevel            ProjectCreationLevelValue  `json:"project_creation_level"`
	AutoDevopsEnabled               bool                       `json:"auto_devops_enabled"`
	SubGroupCreationLevel           SubGroupCreationLevelValue `json:"subgroup_creation_level"`
	EmailsEnabled                   bool                       `json:"emails_enabled"`
	MentionsDisabled                bool                       `json:"mentions_disabled"`
	RunnersToken                    string                     `json:"runners_token"`
	SharedProjects                  []*Project                 `json:"shared_projects"`
	SharedRunnersSetting            SharedRunnersSettingValue  `json:"shared_runners_setting"`
	SharedWithGroups                []struct {
		GroupID          int    `json:"group_id"`
		GroupName        string `json:"group_name"`
		GroupFullPath    string `json:"group_full_path"`
		GroupAccessLevel int    `json:"group_access_level"`
		ExpiresAt        *Date  `json:"expires_at"`
	} `json:"shared_with_groups"`
	LDAPCN                         string             `json:"ldap_cn"`
	LDAPAccess                     AccessLevelValue   `json:"ldap_access"`
	LDAPGroupLinks                 []*LDAPGroupLink   `json:"ldap_group_links"`
	SAMLGroupLinks                 []*SAMLGroupLink   `json:"saml_group_links"`
	SharedRunnersMinutesLimit      int                `json:"shared_runners_minutes_limit"`
	ExtraSharedRunnersMinutesLimit int                `json:"extra_shared_runners_minutes_limit"`
	PreventForkingOutsideGroup     bool               `json:"prevent_forking_outside_group"`
	MarkedForDeletionOn            *Date              `json:"marked_for_deletion_on"`
	CreatedAt                      *time.Time         `json:"created_at"`
	IPRestrictionRanges            string             `json:"ip_restriction_ranges"`
	AllowedEmailDomainsList        string             `json:"allowed_email_domains_list"`
	WikiAccessLevel                AccessControlValue `json:"wiki_access_level"`

	// Deprecated: Use EmailsEnabled instead
	EmailsDisabled bool `json:"emails_disabled"`

	// Deprecated: Use DefaultBranchProtectionDefaults instead
	DefaultBranchProtection int `json:"default_branch_protection"`
}

// BranchProtectionDefaults represents default Git protected branch permissions.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection_defaults
type BranchProtectionDefaults struct {
	AllowedToPush           []*GroupAccessLevel `json:"allowed_to_push,omitempty"`
	AllowForcePush          bool                `json:"allow_force_push,omitempty"`
	AllowedToMerge          []*GroupAccessLevel `json:"allowed_to_merge,omitempty"`
	DeveloperCanInitialPush bool                `json:"developer_can_initial_push,omitempty"`
}

// GroupAccessLevel represents default branch protection defaults access levels.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/groups.html#options-for-default_branch_protection_defaults
type GroupAccessLevel struct {
	AccessLevel *AccessLevelValue `url:"access_level,omitempty" json:"access_level,omitempty"`
}

// LDAPGroupLink represents a GitLab LDAP group link.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/groups.html#ldap-group-links
type LDAPGroupLink struct {
	CN          string           `json:"cn"`
	Filter      string           `json:"filter"`
	GroupAccess AccessLevelValue `json:"group_access"`
	Provider    string           `json:"provider"`
}

// SAMLGroupLink represents a GitLab SAML group link.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/groups.html#saml-group-links
type SAMLGroupLink struct {
	Name         string           `json:"name"`
	AccessLevel  AccessLevelValue `json:"access_level"`
	MemberRoleID int              `json:"member_role_id,omitempty"`
}

// ListGroupsOptions represents the available ListGroups() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/groups.html#list-groups
type ListGroupsOptions struct {
	ListOptions
	SkipGroups           *[]int            `url:"skip_groups,omitempty" del:"," json:"skip_groups,omitempty"`
	AllAvailable         *bool             `url:"all_available,omitempty" json:"all_available,omitempty"`
	Search               *string           `url:"search,omitempty" json:"search,omitempty"`
	OrderBy              *string           `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort                 *string           `url:"sort,omitempty" json:"sort,omitempty"`
	Statistics           *bool             `url:"statistics,omitempty" json:"statistics,omitempty"`
	WithCustomAttributes *bool             `url:"with_custom_attributes,omitempty" json:"with_custom_attributes,omitempty"`
	Owned                *bool             `url:"owned,omitempty" json:"owned,omitempty"`
	MinAccessLevel       *AccessLevelValue `url:"min_access_level,omitempty" json:"min_access_level,omitempty"`
	TopLevelOnly         *bool             `url:"top_level_only,omitempty" json:"top_level_only,omitempty"`
	RepositoryStorage    *string           `url:"repository_storage,omitempty" json:"repository_storage,omitempty"`
}

// ListGroups gets a list of groups (as user: my groups, as admin: all groups).
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/groups.html#list-groups
func (s *GroupsService) ListGroups(ctx context.Context, opts *ListGroupsOptions) (*Records[Group], error) {
	var reply []*Group
	resp, err := s.client.InvokeWithCredential(ctx, http.MethodGet, "groups", opts, &reply)
	if err != nil {
		return nil, err
	}

	return newRecords[Group](opts, reply, resp), nil
}
