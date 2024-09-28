package gitlab

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/nexuer/utils/ptr"
)

func TestTagsService_ListTags(t *testing.T) {
	client := NewClient(testTokenCredential)

	projects, err := client.Projects.ListProjects(context.Background(), &ListProjectsOptions{
		ListOptions: NewListOptions(1, 1),
		OrderBy:     ptr.Ptr("star_count"),
		//Membership:  ptr.Ptr(true),
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(projects) == 0 {
		t.Error(fmt.Errorf("empty projects"))
	}
	project := projects[0]
	t.Logf("project: %s \n", project.WebURL)
	branches, err := client.Tags.ListTags(context.Background(), strconv.Itoa(project.ID), &ListTagsOptions{
		ListOptions: NewListOptions(1, 5),
	})
	if err != nil {
		t.Fatalf("Tags.ListTags returned error: %v", err)
	}
	for _, branch := range branches {
		t.Logf("tag: %s \n", branch.Name)
	}
}
