package gitlab

import (
	"context"
	"testing"

	"github.com/nexuer/utils/ptr"
)

func TestProjectsService_ListProjects(t *testing.T) {
	client := NewClient(testTokenCredential, &Options{Debug: true})

	reply, err := client.Projects.ListProjects(context.Background(), &ListProjectsOptions{
		ListOptions: NewListOptions(1),
		Membership:  ptr.Ptr(true),
		//Search:     ptr.Ptr(""),
		Visibility: ptr.Ptr(PrivateVisibility),
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}

func TestProjectsService_GetProject(t *testing.T) {
	client := NewClient(testTokenCredential, &Options{Debug: true})

	reply, err := client.Projects.GetProject(context.Background(), "1004", &GetProjectOptions{
		Statistics: ptr.Ptr(true),
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}
