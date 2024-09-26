package gitlab

import (
	"context"
	"testing"
)

func TestProjectsService_ListProjects(t *testing.T) {
	client := NewClient(testTokenCredential, &Options{Debug: true})

	// 	查询全部部门
	reply, err := client.Projects.ListProjects(context.Background(), &ListProjectsOptions{
		ListOptions: NewListOptions(1, 10),
		//OrderBy:     ptr.Ptr("last_activity_at"),
		//Membership:  ptr.Ptr(true),
		//Search:      ptr.Ptr(""),
		//Visibility: ptr.Ptr(PrivateVisibility),
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}
