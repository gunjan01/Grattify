package grattify

import (
	"context"

	github "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type client struct {
	ctx context.Context
}

// New returns a new github client.
func (c *client) New() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "XXXXXXXXXXXXXXXXXXXX"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client
}
