package main

import (
	"context"
	"log"

	"github.com/nexuer/go-gitlab"
)

func main() {
	credential := &gitlab.TokenCredential{
		// default endpoint: https://gitlab.com
		//Endpoint: gitlab.CloudEndpoint,
		AccessToken: "token",
	}

	client := gitlab.NewClient(credential, &gitlab.Options{Debug: true})

	ver, err := client.Version.GetVersion(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("version: %+v", ver)

}
