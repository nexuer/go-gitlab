package gitlab

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestBranchesService_ListBranches(t *testing.T) {
	client := NewClient(testTokenCredential)

	projects, _, err := client.Projects.ListProjects(context.Background(), &ListProjectsOptions{
		ListOptions: ListOptions{
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
	branches, pageinfo, err := client.Branches.ListBranches(context.Background(), strconv.Itoa(project.ID), &ListBranchesOptions{
		ListOptions: ListOptions{
			Page:    1,
			PerPage: 20,
		},
	})
	if err != nil {
		t.Fatalf("Branches.ListBranches returned error: %v", err)
	}
	fmt.Printf("pageinfo: %+v\n", pageinfo)
	for _, branch := range branches {
		t.Logf("branch: %s \n", branch.Name)
	}

}
