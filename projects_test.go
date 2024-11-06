package gitlab

import (
	"context"
	"testing"

	"github.com/nexuer/utils/ptr"
)

func TestProjectsService_ListProjects(t *testing.T) {
	client := NewClient(testTokenCredential, &Options{Debug: true})
	opts := &ListProjectsOptions{
		ListOptions:   NewListOptions(1),
		Membership:    ptr.Ptr(true),
		Statistics:    ptr.Ptr(true),
		IncludeHidden: ptr.Ptr(true),
		//Search:     ptr.Ptr(""),
		//Visibility: ptr.Ptr(PrivateVisibility),
	}
	reply, _, err := client.Projects.ListProjects(context.Background(), opts)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v\n", reply)
}

func TestProjectsService_GetProject(t *testing.T) {
	client := NewClient(testTokenCredential, &Options{Debug: true})

	reply, err := client.Projects.GetProject(context.Background(), "971", &GetProjectOptions{
		Statistics: ptr.Ptr(true),
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}

func TestProjectsService_ListWebhooks(t *testing.T) {
	client := NewClient(testTokenCredential, &Options{Debug: true})

	reply, err := client.Projects.ListWebhooks(context.Background(), "971", nil)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}
