package gitlab_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/nexuer/go-gitlab"
)

func TestMembersService_ListGroupMembers(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})

	reply, err := client.Members.ListGroupMembers(context.Background(), "1121", nil)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v\n", reply)
	}

}

func TestMembersService_ListAllProjectMembers(t *testing.T) {
	client := gitlab.NewClient(testTokenCredential, &gitlab.Options{Debug: true})
	reply, err := client.Members.ListAllProjectMembers(context.Background(), "1112", nil)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v\n", reply)
	}

	for _, v := range reply.Records {
		fmt.Println(v.Name, v.AccessLevel.String())
	}
}
