package gitlab

import (
	"context"
	"testing"
	"time"

	"github.com/nexuer/utils/ptr"
)

func TestMilestonesService_ListMilestones(t *testing.T) {
	client := NewClient(testTokenCredential, &Options{Debug: true})

	reply, err := client.Milestones.ListMilestones(context.Background(), "971", &ListMilestonesOptions{
		UpdatedAfter: ptr.Ptr(time.Now()),
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}
