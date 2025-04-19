package gitlab_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/nexuer/go-gitlab"
)

func TestBranchesService_ListBranches(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential)

	projects, err := client.Projects.ListProjects(context.Background(), &gitlab.ListProjectsOptions{
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

	if len(projects.Records) == 0 {
		t.Error(fmt.Errorf("empty projects"))
	}
	project := projects.Records[0]
	t.Logf("project: %s \n", project.WebURL)
	branches, err := client.Branches.ListBranches(context.Background(), strconv.Itoa(project.ID), &gitlab.ListBranchesOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 20,
		},
	})
	if err != nil {
		t.Fatalf("Branches.ListBranches returned error: %v", err)
	}
	fmt.Printf("project: %+v\n", project)
	for _, branch := range branches.Records {
		t.Logf("branch: %s \n", branch.Name)
	}

}
