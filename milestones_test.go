package gitlab_test

import (
	"context"
	"testing"
	"time"

	"github.com/nexuer/go-gitlab"
	"github.com/nexuer/utils/ptr"
)

func TestMilestonesService_ListMilestones(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})

	reply, err := client.Milestones.ListMilestones(context.Background(), "971", &gitlab.ListMilestonesOptions{
		UpdatedAfter: ptr.Ptr(time.Now()),
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}
