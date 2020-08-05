package main

import (
	"fmt"
	"log"
	"net/http"

	"./grattify"
	github "gopkg.in/go-playground/webhooks.v5/github"
)

const (
	path                    = "/webhooks"
	issueAction, pullAction = "opened"
)

func main() {
	hook, _ := github.New(github.Options.Secret(grattify.GithubAccessToken))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.IssuesEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				fmt.Println("Event not found")
			}
		}

		switch payload.(type) {
		case github.IssuesPayload:
			{
				issue := payload.(github.IssuesPayload)
				if issue.Action == issueAction {
					err := grattify.CreateIssueComment(
						issue.Issue.ID,
						issue.Issue.User.Login,
						issue.Repository.Owner.Login,
						issue.Repository.Name,
					)

					if err != nil {
						log.Output(err)
					}
				}
			}

		case github.PullRequestPayload:
			{
				pullPayload := payload.(github.PullRequestPayload)
				if pullPayload.Action == pullAction {
					err := grattify.CreatePRReviewComment(
						pullPayload.PullRequest.User.Login,
						pullPayload.PullRequest.Head.Repo.Owner.Login,
						pullPayload.PullRequest.Head.Repo.Name,
						pullPayload.PullRequest.ID,
					)

					if err != nil {
						log.Output(err)
					}
				}
			}
		}
	})

	http.ListenAndServe(":3000", nil)
}
