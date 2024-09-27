package main

import (
	"context"
	"fmt"
	"os"

	"github.com/nexuer/go-gitlab"
)

func main() {
	credential := &gitlab.TokenCredential{
		// default endpoint: https://gitlab.com
		//Endpoint: gitlab.CloudEndpoint,
		AccessToken: "token",
	}

	client := gitlab.NewClient(credential, &gitlab.Options{Debug: true})

	// 获取版本
	ver, err := client.Version.GetVersion(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("version: %+v\n", ver)

}
