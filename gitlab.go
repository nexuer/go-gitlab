package gitlab

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/nexuer/ghttp"
)

const (
	CloudEndpoint = "https://gitlab.com"
)

type APIVersion string

const (
	APIVersionV4 APIVersion = "v4"
)

type service struct {
	client *Client
}

type Options struct {
	APIVersion APIVersion

	UserAgent string
	Timeout   time.Duration
	Proxy     func(*http.Request) (*url.URL, error)
	Debug     bool
	TLS       *tls.Config
	Limiter   ghttp.Limiter
}

type Client struct {
	cc         *ghttp.Client
	apiVersion APIVersion

	common service

	OAuth *OAuthService
	//
	Branches        *BranchesService
	Commits         *CommitsService
	MergeRequests   *MergeRequestsService
	Tags            *TagsService
	Users           *UsersService
	Projects        *ProjectsService
	Version         *VersionService
	Metadata        *MetadataService
	Releases        *ReleasesService
	RepositoryFiles *RepositoryFilesService
	Milestones      *MilestonesService
	Namespaces      *NamespacesService
	Groups          *GroupsService
	Members         *MembersService
}

func NewClient(credential Credential, opts ...*Options) *Client {
	c := &Client{
		apiVersion: APIVersionV4,
	}

	clientOpts := c.parseOptions(opts...)

	clientOpts = append(clientOpts,
		ghttp.WithNot2xxError(func() error {
			return new(Error)
		}),
	)

	c.cc = ghttp.NewClient(clientOpts...)
	c.common.client = c
	c.OAuth = &OAuthService{client: c.common.client}

	c.Branches = (*BranchesService)(&c.common)
	c.Commits = (*CommitsService)(&c.common)
	c.MergeRequests = (*MergeRequestsService)(&c.common)
	c.Tags = (*TagsService)(&c.common)
	c.Users = (*UsersService)(&c.common)
	c.Projects = (*ProjectsService)(&c.common)
	c.Metadata = (*MetadataService)(&c.common)
	c.Version = (*VersionService)(&c.common)
	c.Releases = (*ReleasesService)(&c.common)
	c.RepositoryFiles = (*RepositoryFilesService)(&c.common)
	c.Milestones = (*MilestonesService)(&c.common)
	c.Namespaces = (*NamespacesService)(&c.common)
	c.Groups = (*GroupsService)(&c.common)
	c.Members = (*MembersService)(&c.common)

	c.SetCredential(credential)
	return c
}

func (c *Client) parseOptions(opts ...*Options) []ghttp.ClientOption {
	var opt *Options
	if len(opts) > 0 && opts[0] != nil {
		opt = opts[0]
	} else {
		opt = new(Options)
	}

	if opt.APIVersion != "" {
		c.apiVersion = opt.APIVersion
	}

	clientOpts := make([]ghttp.ClientOption, 0)

	if opt.UserAgent != "" {
		clientOpts = append(clientOpts, ghttp.WithUserAgent(opt.UserAgent))
	}

	if opt.Debug {
		clientOpts = append(clientOpts, ghttp.WithDebug(opt.Debug))
	}

	if opt.Timeout > 0 {
		clientOpts = append(clientOpts, ghttp.WithTimeout(opt.Timeout))
	}

	if opt.Proxy != nil {
		clientOpts = append(clientOpts, ghttp.WithProxy(opt.Proxy))
	}

	if opt.TLS != nil {
		clientOpts = append(clientOpts, ghttp.WithTLSConfig(opt.TLS))
	}

	if opt.Limiter != nil {
		clientOpts = append(clientOpts, ghttp.WithLimiter(opt.Limiter))
	}

	return clientOpts
}

func (c *Client) SetCredential(credential Credential) {
	var endpoint string
	if credential != nil {
		endpoint = credential.GetEndpoint()
	}

	if endpoint == "" {
		endpoint = CloudEndpoint
	}

	c.cc.SetEndpoint(endpoint)

	if c.OAuth != nil {
		c.OAuth.credential = credential
	}
}

func (c *Client) API(path string, ver ...APIVersion) string {
	if len(ver) > 0 && ver[0] != "" {
		return fmt.Sprintf("/api/%s/%s", ver[0], path)
	}
	return fmt.Sprintf("/api/%s/%s", c.apiVersion, path)
}

func (c *Client) InvokeWithCredential(ctx context.Context, method, path string, args any, reply any, fn ...ghttp.RequestFunc) (*http.Response, error) {
	accessToken, err := c.OAuth.GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	fns := make([]ghttp.RequestFunc, 1, len(fn)+1)
	fns[0] = func(request *http.Request) error {
		return c.OAuth.credential.Auth(request, accessToken)
	}
	fns = append(fns, fn...)
	return c.Invoke(ctx, method, c.API(path), args, reply, fns...)
}

func (c *Client) Invoke(ctx context.Context, method, path string, args any, reply any, fn ...ghttp.RequestFunc) (*http.Response, error) {
	opts := &ghttp.CallOptions{
		BeforeHooks: fn,
	}
	if method == http.MethodGet && args != nil {
		opts.Query = args
		args = nil
	}

	return c.cc.Invoke(ctx, method, path, args, reply, opts)
}

// Error data-validation-and-error-reporting + OAuth error
// GitLab API docs: https://docs.gitlab.com/ee/api/rest/#data-validation-and-error-reporting
// When an attribute is missing, you receive something like:
//
//	{
//	   "message":"400 (Bad request) \"title\" not given"
//	}
//
// When a validation error occurs, error messages are different. They hold all details of validation errors:
//
//	{
//	   "message": {
//	       "<property-name>": [
//	           "<error-message>",
//	           "<error-message>",
//	           ...
//	       ],
//	       "<embed-entity>": {
//	           "<property-name>": [
//	               "<error-message>",
//	               "<error-message>",
//	               ...
//	           ],
//	       }
//	   }
//	}
type Error struct {
	Message          any    `json:"message"`
	Err              string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (e *Error) Error() string {
	if e.ErrorDescription != "" {
		return e.ErrorDescription
	}

	if e.Err != "" {
		return e.Err
	}
	if e.Message != nil {
		switch msg := e.Message.(type) {
		case string:
			return msg
		default:
			b, _ := json.Marshal(e.Message)
			return string(b)
		}
	}
	return ""
}

func IsNotFound(err error) bool {
	code, ok := StatusForErr(err)
	if ok && code == http.StatusNotFound {
		return true
	}
	return false
}

func IsForbidden(err error) bool {
	code, ok := StatusForErr(err)
	if ok && code == http.StatusForbidden {
		return true
	}
	return false
}

func IsUnauthorized(err error) bool {
	code, ok := StatusForErr(err)
	if ok && code == http.StatusUnauthorized {
		return true
	}
	return false
}

func IsTimeout(err error) bool {
	return ghttp.IsTimeout(err)
}

func StatusForErr(err error) (int, bool) {
	return ghttp.StatusForErr(err)
}
