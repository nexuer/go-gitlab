package gitlab

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/zdz1715/ghttp"
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
	UserAgent  string
	APIVersion APIVersion
	Timeout    time.Duration
	Proxy      func(*http.Request) (*url.URL, error)
	Debug      bool
}

type Client struct {
	cc         *ghttp.Client
	apiVersion APIVersion

	common service

	OAuth *OAuthService
	//
	Branches      *BranchesService
	Commits       *CommitsService
	MergeRequests *MergeRequestsService
	Tags          *TagsService
	Users         *UsersService
	Projects      *ProjectsService
	Version       *VersionService
	Metadata      *MetadataService
	Releases      *ReleasesService
}

func NewClient(credential Credential, opts ...*Options) *Client {
	c := &Client{
		apiVersion: APIVersionV4,
	}

	clientOpts := c.parseOptions(opts...)

	clientOpts = append(clientOpts,
		ghttp.WithNot2xxError(func() ghttp.Not2xxError {
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
		clientOpts = append(clientOpts, ghttp.WithDebug(ghttp.DefaultDebug))
	}

	if opt.Timeout > 0 {
		clientOpts = append(clientOpts, ghttp.WithTimeout(opt.Timeout))
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

func (c *Client) api(path string) string {
	return fmt.Sprintf("/api/%s/%s", c.apiVersion, path)
}

func (c *Client) InvokeByCredential(ctx context.Context, method, path string, args any, reply any) error {
	accessToken, err := c.OAuth.GetAccessToken(ctx)
	if err != nil {
		return err
	}

	callOpts := &ghttp.CallOptions{
		BeforeHook: func(request *http.Request) error {
			//request.URL.Path = fmt.Sprintf("api/%s%s", c.apiVersion, request.URL.Path)
			return c.OAuth.credential.Auth(request, accessToken)
		},
	}
	return c.Invoke(ctx, method, c.api(path), args, reply, callOpts)
}

func (c *Client) Invoke(ctx context.Context, method, path string, args any, reply any, callOpts ...*ghttp.CallOptions) error {
	callOpt := &ghttp.CallOptions{}
	if len(callOpts) > 0 && callOpts[0] != nil {
		callOpt = callOpts[0]
	}
	if method == http.MethodGet && args != nil {
		callOpt.Query = args
		args = nil
	}

	_, err := c.cc.Invoke(ctx, method, path, args, reply, callOpt)
	return err
}

// Error data-validation-and-error-reporting + OAuth error
// GitLab API docs: https://docs.gitlab.com/ee/api/rest/#data-validation-and-error-reporting
type Error struct {
	Message any `json:"message"`

	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (e *Error) String() string {
	if e.ErrorDescription != "" {
		return e.ErrorDescription
	}
	if e.Error != "" {
		return e.Error
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

func (e *Error) Reset() {
	e.Message = nil
	e.Error = ""
	e.ErrorDescription = ""
}
