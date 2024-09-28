package gitlab

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/nexuer/utils/ptr"
)

func TestReleasesService_ListReleases(t *testing.T) {
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
	releases, err := client.Releases.ListReleases(context.Background(), strconv.Itoa(project.ID), &ListReleasesOptions{
		ListOptions: &ListOptions{
			Page:    1,
			PerPage: 5,
			OrderBy: "released_at",
			Sort:    SortDesc,
		},
	})
	if err != nil {
		t.Fatalf("Releases.ListReleases returned error: %v", err)
	}
	for _, release := range releases {
		t.Logf("release: %s \n", release.Name)
	}
}
