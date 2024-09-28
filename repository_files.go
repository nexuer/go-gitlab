package gitlab

import (
	"context"
	"fmt"
	"net/http"
)

// RepositoryFilesService handles communication with the repository files
// related methods of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/repository_files.html
type RepositoryFilesService service

// File represents a GitLab repository file.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/repository_files.html
type File struct {
	FileName        string `json:"file_name"`
	FilePath        string `json:"file_path"`
	Size            int    `json:"size"`
	Encoding        string `json:"encoding"`
	Content         string `json:"content"`
	ContentSHA256   string `json:"content_sha256"`
	ExecuteFilemode bool   `json:"execute_filemode"`
	Ref             string `json:"ref"`
	BlobID          string `json:"blob_id"`
	CommitID        string `json:"commit_id"`
	LastCommitID    string `json:"last_commit_id"`
}

// GetFileOptions represents the available GetFile() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/repository_files.html#get-file-from-repository
type GetFileOptions struct {
	Ref *string `query:"ref,omitempty" json:"ref,omitempty"`
}

// GetFile allows you to receive information about a file in repository like
// name, size, content. Note that file content is Base64 encoded.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/repository_files.html#get-file-from-repository
func (s *RepositoryFilesService) GetFile(ctx context.Context, projectID string, filepath string, opts *GetFileOptions) (*File, error) {
	apiEndpoint := fmt.Sprintf("projects/%s/repository/files/%s", projectID, filepath)

	var v File
	if err := s.client.InvokeByCredential(ctx, http.MethodGet, apiEndpoint, opts, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
