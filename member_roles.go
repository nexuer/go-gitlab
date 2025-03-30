package gitlab

// MemberRole represents a GitLab member role.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/member_roles.html
type MemberRole struct {
	ID                         int              `json:"id"`
	Name                       string           `json:"name"`
	Description                string           `json:"description,omitempty"`
	GroupID                    int              `json:"group_id"`
	BaseAccessLevel            AccessLevelValue `json:"base_access_level"`
	AdminCICDVariables         bool             `json:"admin_cicd_variables,omitempty"`
	AdminComplianceFramework   bool             `json:"admin_compliance_framework,omitempty"`
	AdminGroupMembers          bool             `json:"admin_group_member,omitempty"`
	AdminMergeRequests         bool             `json:"admin_merge_request,omitempty"`
	AdminPushRules             bool             `json:"admin_push_rules,omitempty"`
	AdminTerraformState        bool             `json:"admin_terraform_state,omitempty"`
	AdminVulnerability         bool             `json:"admin_vulnerability,omitempty"`
	AdminWebHook               bool             `json:"admin_web_hook,omitempty"`
	ArchiveProject             bool             `json:"archive_project,omitempty"`
	ManageDeployTokens         bool             `json:"manage_deploy_tokens,omitempty"`
	ManageGroupAccesToken      bool             `json:"manage_group_access_tokens,omitempty"`
	ManageMergeRequestSettings bool             `json:"manage_merge_request_settings,omitempty"`
	ManageProjectAccessToken   bool             `json:"manage_project_access_tokens,omitempty"`
	ManageSecurityPolicyLink   bool             `json:"manage_security_policy_link,omitempty"`
	ReadCode                   bool             `json:"read_code,omitempty"`
	ReadRunners                bool             `json:"read_runners,omitempty"`
	ReadDependency             bool             `json:"read_dependency,omitempty"`
	ReadVulnerability          bool             `json:"read_vulnerability,omitempty"`
	RemoveGroup                bool             `json:"remove_group,omitempty"`
	RemoveProject              bool             `json:"remove_project,omitempty"`
}
