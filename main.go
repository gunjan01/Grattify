package main

import (
	"fmt"
	"net/http"

	"./grattify"
	github "gopkg.in/go-playground/webhooks.v5/github"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := github.New(github.Options.Secret("XXXXXXXXXXXXXXXXXXXXX"))

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
				grattify.CreateIssueComment(
					issue.Issue.ID,
					issue.Sender.Login,
					issue.Repository.Owner.Login,
					issue.Repository.Name,
				)
			}

		case github.PullRequestPayload:
			{
				pullPayload := payload.(github.PullRequestPayload)
				if pullPayload.Action == "opened" {
					grattify.CreatePRReviewComment(
						pullPayload.PullRequest.User.Login,
						pullPayload.PullRequest.Head.Repo.Owner.Login,
						pullPayload.PullRequest.Head.Repo.Name,
						pullPayload.PullRequest.ID,
					)
				}
			}
		}
	})

	http.ListenAndServe(":3000", nil)
}
