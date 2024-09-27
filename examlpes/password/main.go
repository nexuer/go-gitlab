package main

import (
	"context"
	"fmt"
	"os"

	"github.com/nexuer/go-gitlab"
)

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
		fmt.Println(err)
		os.Exit(1)
	}
	tk, _ := client.OAuth.GetAccessToken(context.Background())
	fmt.Printf("version: %+v\ntoken: %s\n", ver, tk.AccessToken)
}
