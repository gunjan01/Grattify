package grattify

import (
	"context"
	"fmt"
	"log"
	"os"

	github "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// New returns a new github client.
func New() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GithubAccessToken},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return client
}

// CreateIssueComment creates a comment on the issue.
func CreateIssueComment(id int64, login, owner, repo string) error {
	message := fmt.Sprintf("Thank you for opening an issue @%s. Your contributions are welcome.", login)

	issueComment := github.IssueComment{
		ID:   &id,
		Body: &message,
	}

	githubClient := New()
	comment, _, err := githubClient.Issues.CreateComment(context.Background(), owner, repo, int(id), &issueComment)
	if err != nil {
		return err
	}

	log.SetOutput(os.Stdout)
	log.Print(comment)

	return nil
}

// CreatePRReviewComment creates a thank you comment on the PR.
func CreatePRReviewComment(username string, owner string, repo string, id int64) error {
	message := fmt.Sprintf("Thank you for opening an PR @%s. Your contributions are welcomed ! :)", username)

	pullReqComment := github.PullRequestComment{
		ID:   &id,
		Body: &message,
	}

	githubClient := New()
	comment, _, err := githubClient.PullRequests.CreateComment(context.Background(), owner, repo, int(id), &pullReqComment)
	if err != nil {
		return err
	}

	log.SetOutput(os.Stdout)
	log.Print(comment)

	return nil
}
