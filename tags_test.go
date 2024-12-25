package gitlab_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/nexuer/go-gitlab"
)

func TestTagsService_ListTags(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})

	projects, _, err := client.Projects.ListProjects(context.Background(), &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 1,
			OrderBy: "star_count",
		},
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
	tags, _, err := client.Tags.ListTags(context.Background(), strconv.Itoa(project.ID), &gitlab.ListTagsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 20,
		},
	})
	if err != nil {
		t.Fatalf("Tags.ListTags returned error: %v", err)
	}
	for _, tag := range tags {
		t.Logf("tag: %s create_at: %s\n", tag.Name, tag.CreatedAt)
	}
}
