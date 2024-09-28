package gitlab

import (
	"context"
	"testing"
)

func TestUsersService_ListUsers(t *testing.T) {
	client := NewClient(testTokenCredential, &Options{Debug: true})

	users, err := client.Users.ListUsers(context.Background(), &ListUsersOptions{})
	if err != nil {
		t.Fatalf("Users.ListUsers returned error: %v", err)
	}

	t.Logf("Users.ListUsers returned: %+v", users)
}
