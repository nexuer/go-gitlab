package gitlab_test

import (
	"fmt"
	"net"
	"os"
	"testing"

	"github.com/nexuer/go-gitlab"
)

var testTokenCredential = &gitlab.TokenCredential{
	Endpoint:    os.Getenv("GITLAB_HOST"),
	AccessToken: os.Getenv("GITLAB_TOKEN"),
}

func TestNewListOptions(t *testing.T) {
	ipAddresses, err := net.LookupHost("gitlab.com")
	if err != nil {
		t.Fatal(err)
	}
	for _, ip := range ipAddresses {
		fmt.Println(ip)
	}
}
