package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/nexuer/go-gitlab"
	"os"
)

func main() {
	// docs: https://docs.gitlab.com/ee/api/oauth2.html#authorization-code-flow
	clientID := "6454bdea03a5638793a1a603e431f6d84ef8817555bd73b2a1e12fd35a5f0422"
	clientSecret := "gloas-4573b00d075b3d55e69b22d8695aeceb76b9a4cf8014d4a07608d96fc022a5a8"
	redirectURI := "http://127.0.0.1"
	credential := &gitlab.OAuthCredential{
		Endpoint:     gitlab.CloudEndpoint,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
	}

	client := gitlab.NewClient(credential, &gitlab.Options{Debug: true})

	url := client.OAuth.AuthorizeURL(clientID, redirectURI, "api")

	fmt.Printf("> click url: %s", url)

	codeChan := make(chan string, 1)
	go func() {
		buf := bufio.NewScanner(os.Stdin)
		fmt.Print("\ninput code: ")
		for buf.Scan() {
			codeChan <- buf.Text()
		}
	}()

	select {
	case code := <-codeChan:
		_ = os.Stdin.Close()
		fmt.Printf("auth by code: %+v\n", code)
		token, err := client.OAuth.GetAccessToken(context.Background(), &gitlab.GetAccessTokenOptions{
			Code: code,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("token: %+v\n", token)
	}
}
