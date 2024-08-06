package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-github/v63/github"
	"golang.org/x/oauth2"
	"log"
	"strconv"
	"strings"
	"time"
)

var provisioner Provisioner

func GetProvisioner(provisioner Provisioner) (Provisioner, string, string) {

	var engine string
	var enginedir string

	engine = "terraform"
	log.Println("Default flag is Terraform")
	if engine == "terraform" {
		provisioner = &TerraformProvisioner{}
	} else {
		provisioner = &TofuProvisioner{}
	}

	enginedir = BASETFDIRECTORY

	return provisioner, engine, enginedir
}

// Gets the client to be used by the Activity
func GetGit(git_token string) (gitlclient *github.Client, err error) {

	token := git_token

	log.Printf("the git token passed to the get git function is %s\n", git_token)
	if token == "" {

		log.Panicf("Github Token is not set")
		return nil, errors.New("GitHub token is not set")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client, nil
}

func CreateGitTicket(owner, reponame, token, clustername, subscriptionid, location string, appdeploy bool, appName string) (string, error) {

	log.Printf("the git clustername passed to the get git function is %s\n", clustername)
	log.Printf("the  subscriptionid passed to the get git function is %s\n", subscriptionid)
	log.Printf("the location passed to the git function is %s ", location)

	var gitTitle string
	var gitBody string

	client, err := GetGit(token)

	if appdeploy == true {
		log.Println("We are deploying an app")
		gitTitle = fmt.Sprintf("Deploying App in AKS via Temporal in Region %s", location)
		gitBody = fmt.Sprintf("Please approve the application deployment %s in Azure on Cluster Name is %s with subscription id is  %s\n", appName, clustername, subscriptionid)
	} else {

		log.Printf("Please approve the creation of a cluster in Azure - Cluster Name is %s and the subscription id is  %s\n", clustername, subscriptionid)
		gitTitle = fmt.Sprintf("Provisioning AKS via Temporal in Region %s", location)
		gitBody = fmt.Sprintf("Please approve the creation of a cluster in Azure - Cluster Name is %s and the subscription id is  %s\n", clustername, subscriptionid)
	}
	ctx := context.Background()
	repoOwner := owner
	repoName := reponame
	issueRequest := &github.IssueRequest{
		Title: github.String(gitTitle),
		Body:  github.String(gitBody),
	}

	// Check if the repository exists and you have access
	repo, _, err := client.Repositories.Get(ctx, repoOwner, repoName)
	if err != nil {
		log.Fatalf("Failed to get repository: %v", err)
	}
	fmt.Printf("Repository found: %s\n", repo.GetFullName())

	issue, _, err := client.Issues.Create(ctx, repoOwner, repoName, issueRequest)
	if err != nil {
		log.Println("Failed to create GitHub issue", "Error", err)
		return "", err
	}

	log.Println("Created GitHub issue", "IssueNumber", issue.GetNumber())
	return issue.GetHTMLURL(), nil
}

func PollGitHubIssueStatus(issueURL, gitToken string) error {
	log.Println("Polling GitHub issue status", "IssueURL", issueURL)

	// Authentication

	client, err := GetGit(gitToken)
	//client := github.NewClient(tc)
	ctx := context.Background()
	// Parse issue URL to get repo owner, repo name, and issue number
	// Assuming the issue URL format is https://github.com/{owner}/{repo}/issues/{number}
	issueURLParts := strings.Split(issueURL, "/")
	repoOwner := issueURLParts[3]
	repoName := issueURLParts[4]
	issueNumber, err := strconv.Atoi(issueURLParts[6])
	if err != nil {
		log.Println("Invalid issue number", "Error", err)
		return err
	}

	// Polling loop
	for {
		// Fetch the issue
		issue, _, err := client.Issues.Get(ctx, repoOwner, repoName, issueNumber)
		if err != nil {
			log.Fatalf("Failed to get issue: %v", err)
		}

		// Check if the issue is closed
		if issue.GetState() == "closed" {
			// Fetch the comments on the issue
			comments, _, err := client.Issues.ListComments(ctx, repoOwner, repoName, issueNumber, nil)
			if err != nil {
				log.Fatalf("Failed to get comments: %v", err)
			}

			// Check if any comment contains "approved"
			approved := false
			for _, comment := range comments {
				if strings.Contains(strings.ToLower(comment.GetBody()), "approved") {
					approved = true
					break
				}
			}

			if approved {
				log.Println("GitHub issue is closed and contains approval comment")
				break // Exit the loop
			} else {
				log.Println("GitHub issue is closed but no approval comment found")
				// Optionally, handle the case where the issue is closed but doesn't contain "approved"
			}
		}

		// Sleep for a while before polling again
		time.Sleep(2 * time.Minute) // Adjust the sleep duration as needed

	}

	return nil
}
