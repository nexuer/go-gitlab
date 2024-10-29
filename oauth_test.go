package gitlab

import (
	"fmt"
	"net"
	"os"
	"testing"
)

var testTokenCredential = &TokenCredential{
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
