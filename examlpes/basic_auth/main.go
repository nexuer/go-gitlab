package main

import (
	"context"
	"log"

	"github.com/nexuer/go-gitlab"
)

// This example shows how to create a client with username and password.
func main() {
	// docs: https://docs.gitlab.com/ee/api/oauth2.html#resource-owner-password-credentials-flow
	credential := &gitlab.PasswordCredential{
		// default endpoint: https://gitlab.com
		//Endpoint: gitlab.CloudEndpoint,
		Username: "YourUsername",
		Password: "YourPassword",
	}

	client := gitlab.NewClient(credential, &gitlab.Options{Debug: true})

	ver, err := client.Version.GetVersion(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("version: %+v", ver)
}
