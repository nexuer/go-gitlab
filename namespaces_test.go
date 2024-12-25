package gitlab_test

import (
	"context"
	"testing"

	"github.com/nexuer/go-gitlab"
)

func TestNamespacesService_ListNamespaces(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})

	reply, page, err := client.Namespaces.ListNamespaces(context.Background(), nil)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v\npage: %+v", reply, page)
	}
}
