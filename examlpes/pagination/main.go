package main

import (
	"context"
	"log"

	"github.com/nexuer/go-gitlab"
	"github.com/nexuer/utils/ptr"
)

func main() {
	client := gitlab.NewClient(&gitlab.TokenCredential{
		AccessToken: "glpat-",
	}, &gitlab.Options{Debug: true})

	if err := listAllProjectsByKeySet(client, context.Background()); err != nil {
		log.Fatal(err)
	}
}

func listAllProjectsByKeySet(cc *gitlab.Client, ctx context.Context) error {
	opts := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.NewKeySet("id", gitlab.SortAsc, 1),
		Membership:  ptr.Ptr(true),
	}
	// You can add a retry mechanism here
	for {
		reply, pagination, err := cc.Projects.ListProjects(context.Background(), opts)

		if err != nil {
			return err
		}

		log.Printf("Found project length: %d", len(reply))
		l, ok := pagination.Next()
		if !ok {
			break
		}
		opts.ListOptions = l
	}
	return nil
}

func listAllProjects(cc *gitlab.Client, ctx context.Context) error {
	opts := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.NewListOptions(1, 1),
		Membership:  ptr.Ptr(true),
	}
	// You can add a retry mechanism here
	for {
		reply, pagination, err := cc.Projects.ListProjects(context.Background(), opts)

		if err != nil {
			return err
		}
		log.Printf("Found project length: %d", len(reply))
		if pagination.NextPage == 0 {
			break
		}
		opts.ListOptions.Page = pagination.NextPage
	}
	return nil
}

func listAllProjects1(cc *gitlab.Client, ctx context.Context) error {
	opts := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.NewListOptions(1, 1),
		Membership:  ptr.Ptr(true),
	}
	// You can add a retry mechanism here
	for {
		reply, pagination, err := cc.Projects.ListProjects(context.Background(), opts)

		if err != nil {
			return err
		}
		log.Printf("Found project length: %d", len(reply))
		l, ok := pagination.Next()
		if !ok {
			break
		}
		opts.ListOptions = l
	}
	return nil
}
