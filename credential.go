package gitlab

import (
	"errors"
	"net/http"
)

var (
	ErrCredential = errors.New("invalid credential")
)

type TokenType uint8

const (
	BearerToken TokenType = iota
	PrivateToken
	JobToken
)

type Credential interface {
	GetEndpoint() string
	RequestBody(opts *GetAccessTokenOptions) any
	Auth(req *http.Request, token *AccessToken) error
}

// TokenCredential
// Docs: https://docs.gitlab.com/ee/api/rest/#authentication
type TokenCredential struct {
	Endpoint    string    `json:"endpoint" xml:"endpoint"`
	TokenType   TokenType `json:"type" xml:"type"`
	AccessToken string    `json:"token" xml:"token"`
}

func (t *TokenCredential) GetEndpoint() string {
	return t.Endpoint
}

func (t *TokenCredential) RequestBody(opts *GetAccessTokenOptions) any {
	return nil
}

func (t *TokenCredential) Auth(req *http.Request, token *AccessToken) error {
	if t.AccessToken == "" {
		return errors.New("TokenCredential: no access token")
	}
	switch t.TokenType {
	case JobToken:
		req.Header.Set("JOB-TOKEN", t.AccessToken)
	case PrivateToken:
		req.Header.Set("PRIVATE-TOKEN", t.AccessToken)
	default:
		req.Header.Set("Authorization", "Bearer "+t.AccessToken)
	}
	return nil
}

// OAuthCredential
// docs: https://docs.gitlab.com/ee/api/oauth2.html#authorization-code-flow
type OAuthCredential struct {
	Endpoint     string `json:"endpoint" xml:"endpoint"`
	ClientID     string `json:"client_id" xml:"client_id"`
	ClientSecret string `json:"client_secret" xml:"client_secret"`
	RedirectURI  string `json:"redirect_uri" xml:"redirect_uri"`
}

func (c *OAuthCredential) GetEndpoint() string {
	return c.Endpoint
}

func (c *OAuthCredential) RequestBody(opts *GetAccessTokenOptions) any {
	body := map[string]string{
		"client_id":     c.ClientID,
		"client_secret": c.ClientSecret,
		"redirect_uri":  c.RedirectURI,
	}
	if opts.Code != "" {
		body["grant_type"] = "authorization_code"
		body["code"] = opts.Code
	}

	if opts.RefreshToken != "" {
		body["grant_type"] = "refresh_token"
		body["refresh_token"] = opts.RefreshToken
	}
	return body
}

func (c *OAuthCredential) Auth(req *http.Request, token *AccessToken) error {
	if token != nil {
		req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	}
	return nil
}

// PasswordCredential
// note: The Resource Owner Password Credentials is disabled for users with two-factor authentication turned on.
// These users can access the API using personal access tokens instead.
// docs: https://docs.gitlab.com/ee/api/oauth2.html#resource-owner-password-credentials-flow
type PasswordCredential struct {
	Endpoint string `json:"endpoint" xml:"endpoint"`
	Username string `json:"username" xml:"username"`
	Password string `json:"password" xml:"password"`
}

func (p *PasswordCredential) RequestBody(opts *GetAccessTokenOptions) any {
	return map[string]string{
		"grant_type": "password",
		"username":   p.Username,
		"password":   p.Password,
	}
}

func (p *PasswordCredential) GetEndpoint() string {
	return p.Endpoint
}

func (p *PasswordCredential) Auth(req *http.Request, token *AccessToken) error {
	if token != nil {
		req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	}
	return nil
}
