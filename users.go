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
	CreatedAt *time.Time `json:"created_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	UsageType *string    `json:"usage_type"` // version: 15.7+
}

// AddSSHKey
// GitLab API Docs: https://docs.gitlab.com/ee/api/users.html#add-ssh-key
func (u *UsersService) AddSSHKey(ctx context.Context, req *AddSSHKeyOptions) (*SSHKey, error) {
	const apiEndpoint = "/api/v4/user/keys"
	var key SSHKey
	if err := u.client.InvokeByCredential(ctx, http.MethodPost, apiEndpoint, req, &key); err != nil {
		return nil, err
	}
	return &key, nil
}

// DeleteSSHKey
// GitLab API Docs: https://docs.gitlab.com/ee/api/users.html#delete-ssh-key-for-current-user
func (u *UsersService) DeleteSSHKey(ctx context.Context, keyId string) error {
	apiEndpoint := fmt.Sprintf("/api/v4/user/keys/%s", keyId)
	if err := u.client.InvokeByCredential(ctx, http.MethodDelete, apiEndpoint, nil, nil); err != nil {
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
	LastActivityOn                 *time.Time         `json:"last_activity_on"`
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
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	State     string     `json:"state"`
	CreatedAt *time.Time `json:"created_at"`
	AvatarURL string     `json:"avatar_url"`
	WebURL    string     `json:"web_url"`
}
