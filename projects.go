package gitlab

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// ProjectsService
// GitLab API Docs: https://docs.gitlab.com/ee/api/projects.html
type ProjectsService service

// Project represents a GitLab project.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/projects.html
type Project struct {
	ID                                        int                        `json:"id"`
	Description                               string                     `json:"description"`
	DefaultBranch                             string                     `json:"default_branch"`
	Public                                    bool                       `json:"public"`
	Visibility                                VisibilityValue            `json:"visibility"`
	SSHURLToRepo                              string                     `json:"ssh_url_to_repo"`
	HTTPURLToRepo                             string                     `json:"http_url_to_repo"`
	WebURL                                    string                     `json:"web_url"`
	ReadmeURL                                 string                     `json:"readme_url"`
	TagList                                   []string                   `json:"tag_list"`
	Topics                                    []string                   `json:"topics"`
	Owner                                     *User                      `json:"owner"`
	Name                                      string                     `json:"name"`
	NameWithNamespace                         string                     `json:"name_with_namespace"`
	Path                                      string                     `json:"path"`
	PathWithNamespace                         string                     `json:"path_with_namespace"`
	IssuesEnabled                             bool                       `json:"issues_enabled"`
	OpenIssuesCount                           int                        `json:"open_issues_count"`
	MergeRequestsEnabled                      bool                       `json:"merge_requests_enabled"`
	ApprovalsBeforeMerge                      int                        `json:"approvals_before_merge"`
	JobsEnabled                               bool                       `json:"jobs_enabled"`
	WikiEnabled                               bool                       `json:"wiki_enabled"`
	SnippetsEnabled                           bool                       `json:"snippets_enabled"`
	ResolveOutdatedDiffDiscussions            bool                       `json:"resolve_outdated_diff_discussions"`
	ContainerExpirationPolicy                 *ContainerExpirationPolicy `json:"container_expiration_policy,omitempty"`
	ContainerRegistryEnabled                  bool                       `json:"container_registry_enabled"`
	ContainerRegistryAccessLevel              AccessControlValue         `json:"container_registry_access_level"`
	ContainerRegistryImagePrefix              string                     `json:"container_registry_image_prefix,omitempty"`
	CreatedAt                                 time.Time                  `json:"created_at,omitempty"`
	LastActivityAt                            time.Time                  `json:"last_activity_at,omitempty"`
	CreatorID                                 int                        `json:"creator_id"`
	Namespace                                 *ProjectNamespace          `json:"namespace"`
	Permissions                               *Permissions               `json:"permissions"`
	MarkedForDeletionAt                       *time.Time                 `json:"marked_for_deletion_at"`
	EmptyRepo                                 bool                       `json:"empty_repo"`
	Archived                                  bool                       `json:"archived"`
	AvatarURL                                 string                     `json:"avatar_url"`
	LicenseURL                                string                     `json:"license_url"`
	License                                   *ProjectLicense            `json:"license"`
	SharedRunnersEnabled                      bool                       `json:"shared_runners_enabled"`
	GroupRunnersEnabled                       bool                       `json:"group_runners_enabled"`
	RunnerTokenExpirationInterval             int                        `json:"runner_token_expiration_interval"`
	ForksCount                                int                        `json:"forks_count"`
	StarCount                                 int                        `json:"star_count"`
	RunnersToken                              string                     `json:"runners_token"`
	AllowMergeOnSkippedPipeline               bool                       `json:"allow_merge_on_skipped_pipeline"`
	OnlyAllowMergeIfPipelineSucceeds          bool                       `json:"only_allow_merge_if_pipeline_succeeds"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool                       `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge              bool                       `json:"remove_source_branch_after_merge"`
	PrintingMergeRequestLinkEnabled           bool                       `json:"printing_merge_request_link_enabled"`
	LFSEnabled                                bool                       `json:"lfs_enabled"`
	RepositoryStorage                         string                     `json:"repository_storage"`
	RequestAccessEnabled                      bool                       `json:"request_access_enabled"`
	MergeMethod                               MergeMethodValue           `json:"merge_method"`
	CanCreateMergeRequestIn                   bool                       `json:"can_create_merge_request_in"`
	ForkedFromProject                         *ForkParent                `json:"forked_from_project"`
	Mirror                                    bool                       `json:"mirror"`
	MirrorUserID                              int                        `json:"mirror_user_id"`
	MirrorTriggerBuilds                       bool                       `json:"mirror_trigger_builds"`
	OnlyMirrorProtectedBranches               bool                       `json:"only_mirror_protected_branches"`
	MirrorOverwritesDivergedBranches          bool                       `json:"mirror_overwrites_diverged_branches"`
	PackagesEnabled                           bool                       `json:"packages_enabled"`
	ServiceDeskEnabled                        bool                       `json:"service_desk_enabled"`
	ServiceDeskAddress                        string                     `json:"service_desk_address"`
	IssuesAccessLevel                         AccessControlValue         `json:"issues_access_level"`
	ReleasesAccessLevel                       AccessControlValue         `json:"releases_access_level,omitempty"`
	RepositoryAccessLevel                     AccessControlValue         `json:"repository_access_level"`
	MergeRequestsAccessLevel                  AccessControlValue         `json:"merge_requests_access_level"`
	ForkingAccessLevel                        AccessControlValue         `json:"forking_access_level"`
	WikiAccessLevel                           AccessControlValue         `json:"wiki_access_level"`
	BuildsAccessLevel                         AccessControlValue         `json:"builds_access_level"`
	SnippetsAccessLevel                       AccessControlValue         `json:"snippets_access_level"`
	PagesAccessLevel                          AccessControlValue         `json:"pages_access_level"`
	OperationsAccessLevel                     AccessControlValue         `json:"operations_access_level"`
	AnalyticsAccessLevel                      AccessControlValue         `json:"analytics_access_level"`
	EnvironmentsAccessLevel                   AccessControlValue         `json:"environments_access_level"`
	FeatureFlagsAccessLevel                   AccessControlValue         `json:"feature_flags_access_level"`
	InfrastructureAccessLevel                 AccessControlValue         `json:"infrastructure_access_level"`
	MonitorAccessLevel                        AccessControlValue         `json:"monitor_access_level"`
	AutocloseReferencedIssues                 bool                       `json:"autoclose_referenced_issues"`
	SuggestionCommitMessage                   string                     `json:"suggestion_commit_message"`
	SquashOption                              SquashOptionValue          `json:"squash_option"`
	EnforceAuthChecksOnUploads                bool                       `json:"enforce_auth_checks_on_uploads,omitempty"`
	SharedWithGroups                          []struct {
		GroupID          int    `json:"group_id"`
		GroupName        string `json:"group_name"`
		GroupFullPath    string `json:"group_full_path"`
		GroupAccessLevel int    `json:"group_access_level"`
	} `json:"shared_with_groups"`
	Statistics                               *Statistics        `json:"statistics"`
	Links                                    *Links             `json:"_links,omitempty"`
	ImportURL                                string             `json:"import_url"`
	ImportType                               string             `json:"import_type"`
	ImportStatus                             string             `json:"import_status"`
	ImportError                              string             `json:"import_error"`
	CIDefaultGitDepth                        int                `json:"ci_default_git_depth"`
	CIForwardDeploymentEnabled               bool               `json:"ci_forward_deployment_enabled"`
	CISeperateCache                          bool               `json:"ci_separated_caches"`
	CIJobTokenScopeEnabled                   bool               `json:"ci_job_token_scope_enabled"`
	CIOptInJWT                               bool               `json:"ci_opt_in_jwt"`
	CIAllowForkPipelinesToRunInParentProject bool               `json:"ci_allow_fork_pipelines_to_run_in_parent_project"`
	PublicJobs                               bool               `json:"public_jobs"`
	BuildTimeout                             int                `json:"build_timeout"`
	AutoCancelPendingPipelines               string             `json:"auto_cancel_pending_pipelines"`
	CIConfigPath                             string             `json:"ci_config_path"`
	CustomAttributes                         []*CustomAttribute `json:"custom_attributes"`
	ComplianceFrameworks                     []string           `json:"compliance_frameworks"`
	BuildCoverageRegex                       string             `json:"build_coverage_regex"`
	IssuesTemplate                           string             `json:"issues_template"`
	MergeRequestsTemplate                    string             `json:"merge_requests_template"`
	IssueBranchTemplate                      string             `json:"issue_branch_template"`
	KeepLatestArtifact                       bool               `json:"keep_latest_artifact"`
	MergePipelinesEnabled                    bool               `json:"merge_pipelines_enabled"`
	MergeTrainsEnabled                       bool               `json:"merge_trains_enabled"`
	RestrictUserDefinedVariables             bool               `json:"restrict_user_defined_variables"`
	MergeCommitTemplate                      string             `json:"merge_commit_template"`
	SquashCommitTemplate                     string             `json:"squash_commit_template"`
	AutoDevopsDeployStrategy                 string             `json:"auto_devops_deploy_strategy"`
	AutoDevopsEnabled                        bool               `json:"auto_devops_enabled"`
	BuildGitStrategy                         string             `json:"build_git_strategy"`
	EmailsEnabled                            bool               `json:"emails_enabled"`
	ExternalAuthorizationClassificationLabel string             `json:"external_authorization_classification_label"`
	RequirementsEnabled                      bool               `json:"requirements_enabled"`
	RequirementsAccessLevel                  AccessControlValue `json:"requirements_access_level"`
	SecurityAndComplianceEnabled             bool               `json:"security_and_compliance_enabled"`
	SecurityAndComplianceAccessLevel         AccessControlValue `json:"security_and_compliance_access_level"`
	MergeRequestDefaultTargetSelf            bool               `json:"mr_default_target_self"`

	// Deprecated: Use EmailsEnabled instead
	EmailsDisabled bool `json:"emails_disabled"`
	// Deprecated: This parameter has been renamed to PublicJobs in GitLab 9.0.
	PublicBuilds bool `json:"public_builds"`
}

// Statistics represents a statistics record for a group or project.
type Statistics struct {
	CommitCount           int64 `json:"commit_count"`
	StorageSize           int64 `json:"storage_size"`
	RepositorySize        int64 `json:"repository_size"`
	WikiSize              int64 `json:"wiki_size"`
	LFSObjectsSize        int64 `json:"lfs_objects_size"`
	JobArtifactsSize      int64 `json:"job_artifacts_size"`
	PipelineArtifactsSize int64 `json:"pipeline_artifacts_size"`
	PackagesSize          int64 `json:"packages_size"`
	SnippetsSize          int64 `json:"snippets_size"`
	UploadsSize           int64 `json:"uploads_size"`
}

// Links represents a project web links for self, issues, merge_requests,
// repo_branches, labels, events, members.
type Links struct {
	Self          string `json:"self"`
	Issues        string `json:"issues"`
	MergeRequests string `json:"merge_requests"`
	RepoBranches  string `json:"repo_branches"`
	Labels        string `json:"labels"`
	Events        string `json:"events"`
	Members       string `json:"members"`
	ClusterAgents string `json:"cluster_agents"`
}

// ProjectNamespace represents a project namespace.
type ProjectNamespace struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Kind      string `json:"kind"`
	FullPath  string `json:"full_path"`
	ParentID  int    `json:"parent_id"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

// ForkParent represents the parent project when this is a fork.
type ForkParent struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	NameWithNamespace string `json:"name_with_namespace"`
	Path              string `json:"path"`
	PathWithNamespace string `json:"path_with_namespace"`
	HTTPURLToRepo     string `json:"http_url_to_repo"`
	WebURL            string `json:"web_url"`
	RepositoryStorage string `json:"repository_storage"`
}

// Permissions represents permissions.
type Permissions struct {
	ProjectAccess *ProjectAccess `json:"project_access"`
	GroupAccess   *GroupAccess   `json:"group_access"`
}

// GroupAccess represents group access.
type GroupAccess struct {
	AccessLevel       AccessLevelValue       `json:"access_level"`
	NotificationLevel NotificationLevelValue `json:"notification_level"`
}

// ProjectAccess represents project access.
type ProjectAccess struct {
	AccessLevel       AccessLevelValue       `json:"access_level"`
	NotificationLevel NotificationLevelValue `json:"notification_level"`
}

// ProjectLicense represent the license for a project.
type ProjectLicense struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	HTMLURL   string `json:"html_url"`
	SourceURL string `json:"source_url"`
}

// ContainerExpirationPolicy represents the container expiration policy.
type ContainerExpirationPolicy struct {
	Cadence         string     `json:"cadence"`
	KeepN           int        `json:"keep_n"`
	OlderThan       string     `json:"older_than"`
	NameRegex       string     `json:"name_regex"`
	NameRegexDelete string     `json:"name_regex_delete"`
	NameRegexKeep   string     `json:"name_regex_keep"`
	Enabled         bool       `json:"enabled"`
	NextRunAt       *time.Time `json:"next_run_at"`
}

// ListProjectsOptions represents the available ListProjects() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/projects.html#list-all-projects
type ListProjectsOptions struct {
	*ListOptions `query:",inline"`

	Archived                 *bool             `query:"archived,omitempty"`
	IDAfter                  *int              `query:"id_after,omitempty"`
	IDBefore                 *int              `query:"id_before,omitempty"`
	Imported                 *bool             `query:"imported,omitempty"`
	LastActivityAfter        *time.Time        `query:"last_activity_after,omitempty"`
	LastActivityBefore       *time.Time        `query:"last_activity_before,omitempty"`
	Membership               *bool             `query:"membership,omitempty"`
	MinAccessLevel           *AccessLevelValue `query:"min_access_level,omitempty"`
	OrderBy                  *string           `query:"order_by,omitempty"`
	Owned                    *bool             `query:"owned,omitempty"`
	RepositoryChecksumFailed *bool             `query:"repository_checksum_failed,omitempty"`
	RepositoryStorage        *string           `query:"repository_storage,omitempty"`
	Search                   *string           `query:"search,omitempty"`
	SearchNamespaces         *bool             `query:"search_namespaces,omitempty"`
	Simple                   *bool             `query:"simple,omitempty"`
	Starred                  *bool             `query:"starred,omitempty"`
	Statistics               *bool             `query:"statistics,omitempty"`
	Topic                    *string           `query:"topic,omitempty"`
	Visibility               *VisibilityValue  `query:"visibility,omitempty"`
	WikiChecksumFailed       *bool             `query:"wiki_checksum_failed,omitempty"`
	WithCustomAttributes     *bool             `query:"with_custom_attributes,omitempty"`
	WithIssuesEnabled        *bool             `query:"with_issues_enabled,omitempty"`
	WithMergeRequestsEnabled *bool             `query:"with_merge_requests_enabled,omitempty"`
	WithProgrammingLanguage  *string           `query:"with_programming_language,omitempty"`
}

// ListProjects gets a list of projects accessible by the authenticated user.
// GitLab API docs: https://docs.gitlab.com/ee/api/projects.html#list-all-projects
func (ps *ProjectsService) ListProjects(ctx context.Context, req *ListProjectsOptions) ([]*Project, error) {
	var projects []*Project
	if err := ps.client.InvokeByCredential(ctx, http.MethodGet, "projects", req, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

// GetProjectOptions represents the available GetProject() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/projects.html#get-a-single-project
type GetProjectOptions struct {
	License              *bool `query:"license,omitempty" json:"license,omitempty"`
	Statistics           *bool `query:"statistics,omitempty" json:"statistics,omitempty"`
	WithCustomAttributes *bool `query:"with_custom_attributes,omitempty" json:"with_custom_attributes,omitempty"`
}

// GetProject gets a specific project, identified by project ID or
// NAMESPACE/PROJECT_NAME, which is owned by the authenticated user.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/projects.html#get-a-single-project
func (ps *ProjectsService) GetProject(ctx context.Context, pid string, opts *GetProjectOptions) (*Project, error) {
	u := fmt.Sprintf("projects/%s", pid)
	var project Project
	if err := ps.client.InvokeByCredential(ctx, http.MethodGet, u, opts, &project); err != nil {
		return nil, err
	}
	return &project, nil
}
