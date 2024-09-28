package gitlab

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/nexuer/utils/ptr"
)

func TestBranchesService_ListBranches(t *testing.T) {
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
	branches, err := client.Branches.ListBranches(context.Background(), strconv.Itoa(project.ID), &ListBranchesOptions{
		ListOptions: NewListOptions(1, 5),
	})
	if err != nil {
		t.Fatalf("Branches.ListBranches returned error: %v", err)
	}
	for _, branch := range branches {
		t.Logf("branch: %s \n", branch.Name)
	}
}
