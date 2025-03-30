package gitlab_test

import (
	"context"
	"testing"

	"github.com/nexuer/go-gitlab"
	"github.com/nexuer/utils/ptr"
)

func TestCommitsService_ListCommits(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})

	reply, _, err := client.Commits.ListCommits(context.Background(), "971", &gitlab.ListCommitsOptions{
		RefName:   ptr.Ptr("main"),
		WithStats: ptr.Ptr(true),
		//ListOptions: NewListOptions(1, 10),
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}

}
