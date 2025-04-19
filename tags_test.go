package gitlab_test

import (
	"context"
	"testing"

	"github.com/nexuer/go-gitlab"
)

func TestTagsService_ListTags(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})

	tags, err := client.Tags.ListTags(context.Background(), "1039", &gitlab.ListTagsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 20,
			OrderBy: gitlab.TagsOrderByUpdated,
			Sort:    gitlab.SortDesc,
		},
	})
	if err != nil {
		t.Fatalf("Tags.ListTags returned error: %v", err)
	}
	for _, tag := range tags.Records {
		t.Logf("tag: %s create_at: %s\n", tag.Name, tag.CreatedAt)
	}
}
