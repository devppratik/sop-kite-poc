package main

import (
	"context"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

func getGHReadme(owner, repo, path string) string {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "GH PAT"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	options := github.RepositoryContentGetOptions{}
	content, _, _, err := client.Repositories.GetContents(ctx, owner, repo, path, &options)
	if err != nil {
		panic(err)
	}
	decodedContent, err := content.GetContent()
	if err != nil {
		panic(err)
	}
	return decodedContent
}
