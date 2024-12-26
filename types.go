package gitlab

import (
	"time"
)

// AccessLevelValue represents a permission level within GitLab.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/members.html#roles
type AccessLevelValue int

const (
	NoPermissions            AccessLevelValue = 0
	MinimalAccessPermissions AccessLevelValue = 5
	GuestPermissions         AccessLevelValue = 10
	PlannerPermissions       AccessLevelValue = 15
	ReporterPermissions      AccessLevelValue = 20
	DeveloperPermissions     AccessLevelValue = 30
	MaintainerPermissions    AccessLevelValue = 40
	OwnerPermissions         AccessLevelValue = 50
	AdminPermissions         AccessLevelValue = 60
)

func (a AccessLevelValue) String() string {
	switch a {
	case MinimalAccessPermissions:
		return "MinimalAccess"
	case GuestPermissions:
		return "Guest"
	case PlannerPermissions:
		return "Planner"
	case ReporterPermissions:
		return "Reporter"
	case DeveloperPermissions:
		return "Developer"
	case MaintainerPermissions:
		return "Maintainer"
	case OwnerPermissions:
		return "Owner"
	case AdminPermissions:
		return "Admin"
	default:
		return "NoAccess"
	}
}

type VisibilityValue string

const (
	PrivateVisibility  VisibilityValue = "private"
	InternalVisibility VisibilityValue = "internal"
	PublicVisibility   VisibilityValue = "public"
)

// MergeMethodValue represents a project merge type within GitLab.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/projects.html#project-merge-method
type MergeMethodValue string

const (
	NoFastForwardMerge MergeMethodValue = "merge"
	FastForwardMerge   MergeMethodValue = "ff"
	RebaseMerge        MergeMethodValue = "rebase_merge"
)

// AccessControlValue represents an access control value within GitLab,
// used for managing access to certain project features.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/projects.html
type AccessControlValue string

const (
	DisabledAccessControl AccessControlValue = "disabled"
	EnabledAccessControl  AccessControlValue = "enabled"
	PrivateAccessControl  AccessControlValue = "private"
	PublicAccessControl   AccessControlValue = "public"
)

// NotificationLevelValue represents a notification level.
type NotificationLevelValue int

// SquashOptionValue represents a squash optional level within GitLab.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/projects.html#create-project
type SquashOptionValue string

const (
	SquashOptionNever      SquashOptionValue = "never"
	SquashOptionAlways     SquashOptionValue = "always"
	SquashOptionDefaultOff SquashOptionValue = "default_off"
	SquashOptionDefaultOn  SquashOptionValue = "default_on"
)

type MilestoneState string

const (
	Active MilestoneState = "active"
	Closed MilestoneState = "closed"
)

// BuildStateValue represents a GitLab build state.
type BuildStateValue string

// These constants represent all valid build states.
const (
	Created            BuildStateValue = "created"
	WaitingForResource BuildStateValue = "waiting_for_resource"
	Preparing          BuildStateValue = "preparing"
	Pending            BuildStateValue = "pending"
	Running            BuildStateValue = "running"
	Success            BuildStateValue = "success"
	Failed             BuildStateValue = "failed"
	Canceled           BuildStateValue = "canceled"
	Skipped            BuildStateValue = "skipped"
	Manual             BuildStateValue = "manual"
	Scheduled          BuildStateValue = "scheduled"
)

// Labels is a custom type with specific marshaling characteristics.
type Labels []string

// TasksCompletionStatus represents tasks of the issue/merge request.
type TasksCompletionStatus struct {
	Count          int `json:"count"`
	CompletedCount int `json:"completed_count"`
}

// LinkTypeValue represents a release link type.
type LinkTypeValue string

// List of available release link types.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/releases/links.html#create-a-release-link
const (
	ImageLinkType   LinkTypeValue = "image"
	OtherLinkType   LinkTypeValue = "other"
	PackageLinkType LinkTypeValue = "package"
	RunbookLinkType LinkTypeValue = "runbook"
)

// LinkType is a helper routine that allocates a new LinkType value
// to store v and returns a pointer to it.
func LinkType(v LinkTypeValue) *LinkTypeValue {
	p := new(LinkTypeValue)
	*p = v
	return p
}

type Date struct {
	t time.Time
}

func NewDate(t time.Time) *Date {
	return &Date{t: t}
}

func (d Date) String() string {
	return d.t.Format(time.DateOnly)
}

func (d Date) IsZero() bool {
	return d.t.IsZero()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Date) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	str := string(data)
	if str == "null" {
		return nil
	}
	var err error
	data = data[len(`"`) : len(data)-len(`"`)]
	d.t, err = time.Parse(`"`+time.DateOnly+`"`, str)
	return err
}

// MarshalJSON implements the json.Marshaler interface.
func (d Date) MarshalJSON() ([]byte, error) {
	if d.t.IsZero() {
		return []byte(`null`), nil
	}

	b := make([]byte, 0, len(time.DateOnly)+2)
	b = append(b, '"')
	b = d.t.AppendFormat(b, time.DateOnly)
	b = append(b, '"')

	return b, nil
}

// ProjectCreationLevelValue represents a project creation level within GitLab.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/
type ProjectCreationLevelValue string

// List of available project creation levels.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/
const (
	NoOneProjectCreation      ProjectCreationLevelValue = "noone"
	MaintainerProjectCreation ProjectCreationLevelValue = "maintainer"
	DeveloperProjectCreation  ProjectCreationLevelValue = "developer"
)

// SubGroupCreationLevelValue represents a sub group creation level within GitLab.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/
type SubGroupCreationLevelValue string

// List of available sub group creation levels.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/
const (
	OwnerSubGroupCreationLevelValue      SubGroupCreationLevelValue = "owner"
	MaintainerSubGroupCreationLevelValue SubGroupCreationLevelValue = "maintainer"
)

// SharedRunnersSettingValue determines whether shared runners are enabled for a
// group’s subgroups and projects.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/groups.html#options-for-shared_runners_setting
type SharedRunnersSettingValue string

// List of available shared runner setting levels.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/groups.html#options-for-shared_runners_setting
const (
	EnabledSharedRunnersSettingValue                  SharedRunnersSettingValue = "enabled"
	DisabledAndOverridableSharedRunnersSettingValue   SharedRunnersSettingValue = "disabled_and_overridable"
	DisabledAndUnoverridableSharedRunnersSettingValue SharedRunnersSettingValue = "disabled_and_unoverridable"

	// Deprecated: DisabledWithOverrideSharedRunnersSettingValue is deprecated
	// in favor of DisabledAndOverridableSharedRunnersSettingValue.
	DisabledWithOverrideSharedRunnersSettingValue SharedRunnersSettingValue = "disabled_with_override"
)

const (
	UserNamespaceKind  = "user"
	GroupNamespaceKind = "group"
)
