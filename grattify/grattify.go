package grattify

import (
	"context"
	"fmt"

	github "github.com/google/go-github/github"
)

// createIssue creates an new issue comment.
func CreateIssueComment(id int64, login, owner, repo string) (*github.Issue, error) {
	message := fmt.Sprintf("Thank you for opening an issue @%s", login)
	issueComment := github.IssueComment{
		ID:   &id,
		Body: &message,
	}

	client := client{
		ctx: context.Background(),
	}
	githubClient := client.New()

	_, _, err := githubClient.Issues.CreateComment(client.ctx, owner, repo, int(id), &issueComment)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// CreatePRReviewComment creates a new pull request comment.
func CreatePRReviewComment(username string, owner string, repo string, id int64) (*github.Issue, error) {
	message := fmt.Sprintf("Thank you for opening an PR @%s", username)

	pullReqComment := github.PullRequestComment{
		ID:   &id,
		Body: &message,
	}

	client := client{
		ctx: context.Background(),
	}
	githubClient := client.New()

	_, _, err := githubClient.PullRequests.CreateComment(client.ctx, owner, repo, int(id), &pullReqComment)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

/*
func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "XXXXXXXXXXXXXXX"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", &github.RepositoryListOptions{
		Type:        "private",
		Affiliation: "owner",
	})
	if err == nil {
		for _, repo := range repos {
			fmt.Println(*repo.FullName)
		}
	}

	//createIssue(ctx, client)

}
*/
