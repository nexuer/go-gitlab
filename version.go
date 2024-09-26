package gitlab

import (
	"context"
	"net/http"
)

// VersionService
// GitLab API docs: https://docs.gitlab.com/ee/api/version.html
type VersionService service

type Version struct {
	Version  string `json:"version"`
	Revision string `json:"revision"`
}

func (vs *VersionService) GetVersion(ctx context.Context) (*Version, error) {
	const apiEndpoint = "/api/v4/version"
	var v Version
	if err := vs.client.InvokeByCredential(ctx, http.MethodGet, apiEndpoint, nil, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
