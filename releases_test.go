package gitlab_test

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"testing"

	"github.com/nexuer/go-gitlab"
)

func TestReleasesService_ListReleases(t *testing.T) {
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
	releases, err := client.Releases.ListReleases(context.Background(), strconv.Itoa(project.ID), &gitlab.ListReleasesOptions{
		//ListOptions: NewKeySet("", SortAsc),
	})
	if err != nil {
		t.Fatalf("Releases.ListReleases returned error: %v", err)
	}
	for _, release := range releases.Records {
		t.Logf("release: %s \n", release.Name)
	}
}

func TestReleasesService_GetRelease(t *testing.T) {
	fmt.Println(url.QueryEscape("diaspora/diaspora"))
}
