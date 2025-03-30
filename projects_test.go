package gitlab_test

import (
	"context"
	"testing"

	"github.com/nexuer/go-gitlab"
	"github.com/nexuer/utils/ptr"
)

func TestProjectsService_ListProjects(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})
	opts := &gitlab.ListProjectsOptions{
		ListOptions:   gitlab.NewListOptions(1),
		Membership:    ptr.Ptr(true),
		Statistics:    ptr.Ptr(true),
		IncludeHidden: ptr.Ptr(true),
		//Search:     ptr.Ptr(""),
		//Visibility: ptr.Ptr(PrivateVisibility),
	}
	reply, page, err := client.Projects.ListProjects(context.Background(), opts)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("page: %+v\n%v\n", page, reply)
}

func TestProjectsService_GetProject(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})

	reply, err := client.Projects.GetProject(context.Background(), "971", &gitlab.GetProjectOptions{
		Statistics: ptr.Ptr(true),
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}

func TestProjectsService_ListWebhooks(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})

	reply, err := client.Projects.ListWebhooks(context.Background(), "971", nil)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}
