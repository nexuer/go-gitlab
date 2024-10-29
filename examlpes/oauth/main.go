package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/nexuer/go-gitlab"
)

func main() {
	// docs: https://docs.gitlab.com/ee/api/oauth2.html#authorization-code-flow
	clientID := "YourClientID"
	clientSecret := "YourClientSecret"
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

		// Fetch version
		ver, err := client.Version.GetVersion(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("version: %+v\n", ver)

		// Refresh token
		token, err = client.OAuth.GetAccessToken(context.Background(), &gitlab.GetAccessTokenOptions{
			RefreshToken: token.RefreshToken,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("RefreshToken: %+v\n", token)
	}
}
