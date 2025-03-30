package gitlab

import (
	"context"
	"testing"

	"github.com/nexuer/utils/ptr"
)

func TestUsersService_ListUsers(t *testing.T) {
	client := NewClient(testTokenCredential, &Options{Debug: true})

	users, _, err := client.Users.ListUsers(context.Background(), &ListUsersOptions{
		ListOptions:   NewKeySet("username", SortAsc),
		ExcludeHumans: ptr.Ptr(true),
	})
	if err != nil {
		t.Fatalf("Users.ListUsers returned error: %v", err)
	}

	t.Logf("Users.ListUsers returned: %+v", users)
}

func TestUsersService_ListSSHKeys(t *testing.T) {
	client := NewClient(testTokenCredential, &Options{Debug: true})

	keys, err := client.Users.ListSSHKeys(context.Background())
	if err != nil {
		t.Fatalf("Users.ListSSHKeys returned error: %v", err)
	}

	t.Logf("Users.ListSSHKeys returned: %+v", keys)
}
