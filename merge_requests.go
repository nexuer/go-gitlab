package gitlab

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// MergeRequestsService handles communication with the merge requests related
// methods of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/merge_requests.html
type MergeRequestsService service

type MergeRequest struct {
	ID                        int                 `json:"id"`
	IID                       int                 `json:"iid"`
	TargetBranch              string              `json:"target_branch"`
	SourceBranch              string              `json:"source_branch"`
	ProjectID                 int                 `json:"project_id"`
	Title                     string              `json:"title"`
	State                     string              `json:"state"`
	CreatedAt                 time.Time           `json:"created_at"`
	UpdatedAt                 time.Time           `json:"updated_at"`
	Upvotes                   int                 `json:"upvotes"`
	Downvotes                 int                 `json:"downvotes"`
	Author                    *BasicUser          `json:"author"`
	Assignee                  *BasicUser          `json:"assignee"`
	Assignees                 []*BasicUser        `json:"assignees"`
	Reviewers                 []*BasicUser        `json:"reviewers"`
	SourceProjectID           int                 `json:"source_project_id"`
	TargetProjectID           int                 `json:"target_project_id"`
	Labels                    Labels              `json:"labels"`
	LabelDetails              []*LabelDetails     `json:"label_details"`
	Description               string              `json:"description"`
	Draft                     bool                `json:"draft"`
	WorkInProgress            bool                `json:"work_in_progress"`
	Milestone                 *Milestone          `json:"milestone"`
	MergeWhenPipelineSucceeds bool                `json:"merge_when_pipeline_succeeds"`
	DetailedMergeStatus       string              `json:"detailed_merge_status"`
	MergeError                string              `json:"merge_error"`
	MergedBy                  *BasicUser          `json:"merged_by"`
	MergedAt                  *time.Time          `json:"merged_at"`
	ClosedBy                  *BasicUser          `json:"closed_by"`
	ClosedAt                  *time.Time          `json:"closed_at"`
	Subscribed                bool                `json:"subscribed"`
	SHA                       string              `json:"sha"`
	MergeCommitSHA            string              `json:"merge_commit_sha"`
	SquashCommitSHA           string              `json:"squash_commit_sha"`
	UserNotesCount            int                 `json:"user_notes_count"`
	ChangesCount              string              `json:"changes_count"`
	ShouldRemoveSourceBranch  bool                `json:"should_remove_source_branch"`
	ForceRemoveSourceBranch   bool                `json:"force_remove_source_branch"`
	AllowCollaboration        bool                `json:"allow_collaboration"`
	WebURL                    string              `json:"web_url"`
	References                *IssueReferences    `json:"references"`
	DiscussionLocked          bool                `json:"discussion_locked"`
	Changes                   []*MergeRequestDiff `json:"changes"`
	User                      struct {
		CanMerge bool `json:"can_merge"`
	} `json:"user"`
	TimeStats    *TimeStats    `json:"time_stats"`
	Squash       bool          `json:"squash"`
	Pipeline     *PipelineInfo `json:"pipeline"`
	HeadPipeline *Pipeline     `json:"head_pipeline"`
	DiffRefs     struct {
		BaseSha  string `json:"base_sha"`
		HeadSha  string `json:"head_sha"`
		StartSha string `json:"start_sha"`
	} `json:"diff_refs"`
	DivergedCommitsCount        int                    `json:"diverged_commits_count"`
	RebaseInProgress            bool                   `json:"rebase_in_progress"`
	ApprovalsBeforeMerge        int                    `json:"approvals_before_merge"`
	Reference                   string                 `json:"reference"`
	FirstContribution           bool                   `json:"first_contribution"`
	TaskCompletionStatus        *TasksCompletionStatus `json:"task_completion_status"`
	HasConflicts                bool                   `json:"has_conflicts"`
	BlockingDiscussionsResolved bool                   `json:"blocking_discussions_resolved"`
	Overflow                    bool                   `json:"overflow"`

	// Deprecated: This parameter is replaced by DetailedMergeStatus in GitLab 15.6.
	MergeStatus string `json:"merge_status"`
}

// MergeRequestDiff represents Gitlab merge request diff.
//
// Gitlab API docs:
// https://docs.gitlab.com/ee/api/merge_requests.html#list-merge-request-diffs
type MergeRequestDiff struct {
	OldPath     string `json:"old_path"`
	NewPath     string `json:"new_path"`
	AMode       string `json:"a_mode"`
	BMode       string `json:"b_mode"`
	Diff        string `json:"diff"`
	NewFile     bool   `json:"new_file"`
	RenamedFile bool   `json:"renamed_file"`
	DeletedFile bool   `json:"deleted_file"`
}

// CreateMergeRequestOptions represents the available CreateMergeRequest()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/merge_requests.html#create-mr
type CreateMergeRequestOptions struct {
	Title              *string `json:"title,omitempty" query:"title"`
	Description        *string `json:"description,omitempty" query:"description"`
	SourceBranch       *string `json:"source_branch,omitempty" query:"source_branch"`
	TargetBranch       *string `json:"target_branch,omitempty" query:"target_branch"`
	Labels             *Labels `json:"labels,omitempty" query:"labels"`
	AssigneeID         *int    `json:"assignee_id,omitempty" query:"assignee_id"`
	AssigneeIDs        []int   `json:"assignee_i_ds,omitempty" query:"assignee_i_ds"`
	ReviewerIDs        []int   `json:"reviewer_i_ds,omitempty" query:"reviewer_i_ds"`
	TargetProjectID    *int    `json:"target_project_id,omitempty" query:"target_project_id"`
	MilestoneID        *int    `json:"milestone_id,omitempty" query:"milestone_id"`
	RemoveSourceBranch *bool   `json:"remove_source_branch,omitempty" query:"remove_source_branch"`
	Squash             *bool   `json:"squash,omitempty" query:"squash"`
	AllowCollaboration *bool   `json:"allow_collaboration,omitempty" query:"allow_collaboration"`
}

func (s *MergeRequestsService) CreateMergeRequest(ctx context.Context, projectId string, opts *CreateMergeRequestOptions) (*MergeRequest, error) {
	apiEndpoint := fmt.Sprintf("projects/%s/merge_requests", projectId)
	var v *MergeRequest
	if err := s.client.InvokeByCredential(ctx, http.MethodPost, apiEndpoint, opts, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// AcceptMergeRequestOptions represents the available AcceptMergeRequest()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/merge_requests.html#merge-a-merge-request
type AcceptMergeRequestOptions struct {
	MergeCommitMessage        *string `json:"merge_commit_message,omitempty" query:"merge_commit_message"`
	MergeWhenPipelineSucceeds *bool   `json:"merge_when_pipeline_succeeds,omitempty" query:"merge_when_pipeline_succeeds"`
	SHA                       *string `json:"sha,omitempty" query:"sha"`
	ShouldRemoveSourceBranch  *bool   `json:"should_remove_source_branch,omitempty" query:"should_remove_source_branch"`
	SquashCommitMessage       *string `json:"squash_commit_message,omitempty" query:"squash_commit_message"`
	Squash                    *bool   `json:"squash,omitempty" query:"squash"`
}

// AcceptMergeRequest
// 401	Unauthorized	This user does not have permission to accept this merge request.
// 405	Method Not Allowed	The merge request is not able to be merged.
// 409	SHA does not match HEAD of source branch	The provided sha parameter does not match the HEAD of the source.
// 422	Branch cannot be merged	The merge request failed to merge.
func (s *MergeRequestsService) AcceptMergeRequest(ctx context.Context, projectId string, iid int, opts *AcceptMergeRequestOptions) (*MergeRequest, error) {
	apiEndpoint := fmt.Sprintf("projects/%s/merge_requests/%d/merge", projectId, iid)
	var v *MergeRequest
	if err := s.client.InvokeByCredential(ctx, http.MethodPut, apiEndpoint, opts, &v); err != nil {
		return nil, err
	}
	return v, nil
}
