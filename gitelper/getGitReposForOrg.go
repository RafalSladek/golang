package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("as24_git_token")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	var allRepos []*github.Repository
	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.ListByOrg(ctx, "autoscout24", opt)

	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
	allRepos = append(allRepos, repos...)

	fmt.Printf("\n%v\n", github.Stringify(repos))
}
