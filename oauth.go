package gitlab

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// OAuthService
// GitLab API Docs: https://docs.gitlab.com/ee/api/oauth2.html
type OAuthService struct {
	client     *Client
	credential Credential
	store      store
}

type store struct {
	val    *AccessToken
	expire time.Time
}

func (s *store) value() *AccessToken {
	if s.isExpired() {
		return nil
	}
	s.val.ExpiresIn = int64(s.expire.Sub(time.Now()).Seconds())
	return s.val
}

func (s *store) isExpired() bool {
	if s.val == nil || s.val.AccessToken == "" {
		return true
	}
	return time.Now().After(s.expire)
}

func (s *store) memory(at *AccessToken) {
	s.val = at
	if at != nil {
		// expire 30 seconds in advance to avoid network delays
		s.expire = time.Now().Add(time.Duration(at.ExpiresIn-30) * time.Second)
	}
}

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int64  `json:"expires_in"`
	CreatedAt    int64  `json:"created_at"`
}

func (oa *OAuthService) AuthorizeURL(clientId, redirectUri, scope string) string {
	u := ""
	if oa.credential != nil {
		u = oa.credential.GetEndpoint()
	}
	return fmt.Sprintf("%s/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s",
		u,
		clientId,
		url.QueryEscape(redirectUri),
		url.QueryEscape(scope),
	)
}

type GetAccessTokenOptions struct {
	Code         string
	RefreshToken string
}

func (oa *OAuthService) GetAccessToken(ctx context.Context, opts ...*GetAccessTokenOptions) (*AccessToken, error) {
	opt := &GetAccessTokenOptions{}
	if len(opts) > 0 && opts[0] != nil {
		opt = opts[0]
	}

	if opt.RefreshToken == "" {
		storeToken := oa.store.value()
		if storeToken != nil {
			return storeToken, nil
		}
	}

	if oa.credential == nil {
		return nil, ErrCredential
	}

	req := oa.credential.RequestBody(opt)
	if req == nil {
		return nil, nil
	}

	var respBody AccessToken
	if _, err := oa.client.Invoke(ctx, http.MethodPost, "/oauth/token", req, &respBody); err != nil {
		return nil, err
	}
	oa.store.memory(&respBody)
	return &respBody, nil
}
