package gitlab_test

import (
	"context"
	"testing"

	"github.com/nexuer/go-gitlab"
)

func TestGroupsService_ListGroups(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})

	reply, err := client.Groups.ListGroups(context.Background(), nil)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v\n", reply)
	}
}
