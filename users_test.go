package gitlab_test

import (
	"context"
	"testing"

	"github.com/nexuer/go-gitlab"
)

func TestUsersService_ListUsers(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})

	users, _, err := client.Users.ListUsers(context.Background(), &gitlab.ListUsersOptions{
		ListOptions: gitlab.NewKeySet("username", gitlab.SortAsc),
	})
	if err != nil {
		t.Fatalf("Users.ListUsers returned error: %v", err)
	}

	t.Logf("Users.ListUsers returned: %+v", users)
}

func TestUsersService_ListSSHKeys(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})

	keys, err := client.Users.ListSSHKeys(context.Background())
	if err != nil {
		t.Fatalf("Users.ListSSHKeys returned error: %v", err)
	}

	t.Logf("Users.ListSSHKeys returned: %+v", keys)
}
