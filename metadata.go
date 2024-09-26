package gitlab

import (
	"context"
	"net/http"
)

// MetadataService
// GitLab API docs: https://docs.gitlab.com/ee/api/metadata.html
type MetadataService service

// Metadata represents a GitLab instance version.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/metadata.html
type Metadata struct {
	Version  string `json:"version"`
	Revision string `json:"revision"`
	KAS      struct {
		Enabled     bool   `json:"enabled"`
		ExternalURL string `json:"externalUrl"`
		Version     string `json:"version"`
	} `json:"kas"`
	Enterprise bool `json:"enterprise"`
}

func (ms *MetadataService) GetMetadata(ctx context.Context) (*Metadata, error) {
	const apiEndpoint = "/api/v4/metadata"
	var v Metadata
	if err := ms.client.InvokeByCredential(ctx, http.MethodGet, apiEndpoint, nil, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
