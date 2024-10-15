package gitlab

import (
	"context"
	"testing"

	"github.com/nexuer/utils/ptr"
)

func TestCommitsService_ListCommits(t *testing.T) {
	client := NewClient(testTokenCredential, &Options{Debug: true})

	reply, err := client.Commits.ListCommits(context.Background(), "971", &ListCommitsOptions{
		RefName: ptr.Ptr("main"),
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}
