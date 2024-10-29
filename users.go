package gitlab

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

// UsersService
// GitLab API Docs: https://docs.gitlab.com/ee/api/users.html
type UsersService service

type AddSSHKeyOptions struct {
	Key       *string `json:"key,omitempty"`
	Title     *string `json:"title,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
	UsageType *string `json:"usage_type,omitempty"` // version: 15.7+
}

type SSHKey struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Key       string     `json:"key"`
	CreatedAt time.Time  `json:"created_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	UsageType *string    `json:"usage_type"` // version: 15.7+
}

// ListSSHKeys
// GitLab API Docs: https://docs.gitlab.com/ee/api/user_keys.html#list-your-ssh-keys
func (u *UsersService) ListSSHKeys(ctx context.Context) ([]*SSHKey, error) {
	const apiEndpoint = "user/keys"
	var keys []*SSHKey
	if _, err := u.client.InvokeByCredential(ctx, http.MethodGet, apiEndpoint, nil, &keys); err != nil {
		return nil, err
	}
	return keys, nil
}

// AddSSHKey
// GitLab API Docs: https://docs.gitlab.com/ee/api/user_keys.html#add-an-ssh-key-to-your-account
func (u *UsersService) AddSSHKey(ctx context.Context, req *AddSSHKeyOptions) (*SSHKey, error) {
	const apiEndpoint = "user/keys"
	var key SSHKey
	if _, err := u.client.InvokeByCredential(ctx, http.MethodPost, apiEndpoint, req, &key); err != nil {
		return nil, err
	}
	return &key, nil
}

// DeleteSSHKey
// GitLab API Docs: https://docs.gitlab.com/ee/api/user_keys.html#delete-an-ssh-key-from-your-account
func (u *UsersService) DeleteSSHKey(ctx context.Context, keyId string) error {
	apiEndpoint := fmt.Sprintf("user/keys/%s", keyId)
	if _, err := u.client.InvokeByCredential(ctx, http.MethodDelete, apiEndpoint, nil, nil); err != nil {
		return err
	}
	return nil
}

// User represents a GitLab user.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/users.html
type User struct {
	ID                             int                `json:"id"`
	Username                       string             `json:"username"`
	Email                          string             `json:"email"`
	Name                           string             `json:"name"`
	State                          string             `json:"state"`
	Locked                         bool               `json:"locked"`
	WebURL                         string             `json:"web_url"`
	CreatedAt                      *time.Time         `json:"created_at"`
	Bio                            string             `json:"bio"`
	Bot                            bool               `json:"bot"`
	Location                       string             `json:"location"`
	PublicEmail                    string             `json:"public_email"`
	Skype                          string             `json:"skype"`
	Linkedin                       string             `json:"linkedin"`
	Twitter                        string             `json:"twitter"`
	WebsiteURL                     string             `json:"website_url"`
	Organization                   string             `json:"organization"`
	JobTitle                       string             `json:"job_title"`
	ExternUID                      string             `json:"extern_uid"`
	Provider                       string             `json:"provider"`
	ThemeID                        int                `json:"theme_id"`
	LastActivityOn                 *Date              `json:"last_activity_on"`
	ColorSchemeID                  int                `json:"color_scheme_id"`
	IsAdmin                        bool               `json:"is_admin"`
	AvatarURL                      string             `json:"avatar_url"`
	CanCreateGroup                 bool               `json:"can_create_group"`
	CanCreateProject               bool               `json:"can_create_project"`
	ProjectsLimit                  int                `json:"projects_limit"`
	CurrentSignInAt                *time.Time         `json:"current_sign_in_at"`
	CurrentSignInIP                *net.IP            `json:"current_sign_in_ip"`
	LastSignInAt                   *time.Time         `json:"last_sign_in_at"`
	LastSignInIP                   *net.IP            `json:"last_sign_in_ip"`
	ConfirmedAt                    *time.Time         `json:"confirmed_at"`
	TwoFactorEnabled               bool               `json:"two_factor_enabled"`
	Note                           string             `json:"note"`
	Identities                     []*UserIdentity    `json:"identities"`
	External                       bool               `json:"external"`
	PrivateProfile                 bool               `json:"private_profile"`
	SharedRunnersMinutesLimit      int                `json:"shared_runners_minutes_limit"`
	ExtraSharedRunnersMinutesLimit int                `json:"extra_shared_runners_minutes_limit"`
	UsingLicenseSeat               bool               `json:"using_license_seat"`
	CustomAttributes               []*CustomAttribute `json:"custom_attributes"`
	NamespaceID                    int                `json:"namespace_id"`
}

// UserIdentity represents a user identity.
type UserIdentity struct {
	Provider  string `json:"provider"`
	ExternUID string `json:"extern_uid"`
}

// BasicUser included in other service responses (such as merge requests, pipelines, etc).
type BasicUser struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	AvatarURL string    `json:"avatar_url"`
	WebURL    string    `json:"web_url"`
}

// ListUsersOptions represents the available ListUsers() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/users.html#list-users
type ListUsersOptions struct {
	ListOptions `query:",inline"`

	Active          *bool `query:"active,omitempty" json:"active,omitempty"`
	Auditors        *bool `query:"auditors,omitempty" json:"auditors,omitempty"`
	Blocked         *bool `query:"blocked,omitempty" json:"blocked,omitempty"`
	ExcludeInternal *bool `query:"exclude_internal,omitempty" json:"exclude_internal,omitempty"`
	ExcludeExternal *bool `query:"exclude_external,omitempty" json:"exclude_external,omitempty"`
	ExcludeActive   *bool `query:"exclude_active,omitempty" json:"exclude_active,omitempty"`
	ExcludeHumans   *bool `query:"exclude_humans,omitempty" json:"exclude_humans,omitempty"`

	// The options below are only available for admins.
	Search               *string    `query:"search,omitempty" json:"search,omitempty"`
	Username             *string    `query:"username,omitempty" json:"username,omitempty"`
	ExternalUID          *string    `query:"extern_uid,omitempty" json:"extern_uid,omitempty"`
	Provider             *string    `query:"provider,omitempty" json:"provider,omitempty"`
	CreatedBefore        *time.Time `query:"created_before,omitempty" json:"created_before,omitempty"`
	CreatedAfter         *time.Time `query:"created_after,omitempty" json:"created_after,omitempty"`
	OrderBy              *string    `query:"order_by,omitempty" json:"order_by,omitempty"`
	TwoFactor            *string    `query:"two_factor,omitempty" json:"two_factor,omitempty"`
	Admins               *bool      `query:"admins,omitempty" json:"admins,omitempty"`
	External             *bool      `query:"external,omitempty" json:"external,omitempty"`
	Humans               *bool      `query:"humans,omitempty" json:"humans,omitempty"`
	SkipLdap             *bool      `query:"skip_ldap,omitempty" json:"skip_ldap,omitempty"`
	WithoutProjects      *bool      `query:"without_projects,omitempty" json:"without_projects,omitempty"`
	WithCustomAttributes *bool      `query:"with_custom_attributes,omitempty" json:"with_custom_attributes,omitempty"`
	WithoutProjectBots   *bool      `query:"without_project_bots,omitempty" json:"without_project_bots,omitempty"`
}

// ListUsers gets a list of users.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/users.html#list-users
func (u *UsersService) ListUsers(ctx context.Context, opts *ListUsersOptions) ([]*User, error) {
	var users []*User
	if _, err := u.client.InvokeByCredential(ctx, http.MethodGet, "users", opts, &users); err != nil {
		return nil, err
	}
	return users, nil
}
