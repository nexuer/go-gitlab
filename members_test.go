package gitlab_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/nexuer/go-gitlab"
)

func TestMembersService_ListGroupMembers(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})

	reply, page, err := client.Members.ListGroupMembers(context.Background(), "1121", nil)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v\npage: %+v", reply, page)
	}

}

func TestMembersService_ListAllProjectMembers(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})
	reply, page, err := client.Members.ListAllProjectMembers(context.Background(), "1112", nil)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v\npage: %+v", reply, page)
	}

	for _, v := range reply {
		fmt.Println(v.Name, v.AccessLevel.String())
	}
}
