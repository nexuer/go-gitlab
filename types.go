package gitlab

// AccessLevelValue represents a permission level within GitLab.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/members.html#roles
type AccessLevelValue int

const (
	NoPermissions            AccessLevelValue = 0
	MinimalAccessPermissions AccessLevelValue = 5
	GuestPermissions         AccessLevelValue = 10
	ReporterPermissions      AccessLevelValue = 20
	DeveloperPermissions     AccessLevelValue = 30
	MaintainerPermissions    AccessLevelValue = 40
	OwnerPermissions         AccessLevelValue = 50
)

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
