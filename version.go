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
	const apiEndpoint = "version"
	var v Version
	if _, err := vs.client.InvokeWithCredential(ctx, http.MethodGet, apiEndpoint, nil, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
