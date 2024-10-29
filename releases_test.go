package gitlab

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"testing"
)

func TestReleasesService_ListReleases(t *testing.T) {
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
	releases, err := client.Releases.ListReleases(context.Background(), strconv.Itoa(project.ID), &ListReleasesOptions{
		//ListOptions: NewKeySet("", SortAsc),
	})
	if err != nil {
		t.Fatalf("Releases.ListReleases returned error: %v", err)
	}
	for _, release := range releases {
		t.Logf("release: %s \n", release.Name)
	}
}

func TestReleasesService_GetRelease(t *testing.T) {
	fmt.Println(url.QueryEscape("diaspora/diaspora"))
}
