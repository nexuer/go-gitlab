package gitlab

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/nexuer/utils/ptr"
)

func TestRepositoryFilesService_GetFile(t *testing.T) {
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

	file, err := client.RepositoryFiles.GetFile(context.Background(), strconv.Itoa(project.ID), ".gitignore", &GetFileOptions{
		Ref: ptr.Ptr("master"),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Branches.ListBranches returned: %+v", file)

	bs, err := file.GetContent()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("file content: \n%s", string(bs))
}
