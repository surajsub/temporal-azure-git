package activities

import (
	"context"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/activity"
)

func (a *AKSImpl) GitCreateTicketActivity(ctx context.Context, owner, reponame, token, clustername, subid, location, appName string) (string, error) {

	ticket, err := utils.CreateGitTicket(owner, reponame, token, clustername, subid, location, false, appName)
	if err != nil {
		return "", err
	}

	activity.GetLogger(ctx).Info(utils.INIT_ACTIVITY, ticket)
	return ticket, nil
}

func (a *AKSImpl) GitPollTicketActivity(ctx context.Context, ticketid, token string) error {
	err := utils.PollGitHubIssueStatus(ticketid, token)
	if err != nil {
		return err
	}

	activity.GetLogger(ctx).Info("Executing the Git Poll Activity")
	return nil
}

func (a *AKSImpl) GitCreateTicketAppDeployActivity(ctx context.Context, owner, reponame, token, clustername, subid, location, appName string) (string, error) {

	ticket, err := utils.CreateGitTicket(owner, reponame, token, clustername, subid, location, true, appName)
	if err != nil {
		return "", err
	}

	activity.GetLogger(ctx).Info(utils.INIT_ACTIVITY, ticket)
	return ticket, nil
}
