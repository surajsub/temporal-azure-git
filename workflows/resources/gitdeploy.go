package resources

import (
	"github.com/surajsub/temporal-azure-git/activities"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/workflow"
	"time"
)

func GitAppDeployWorkflow(ctx workflow.Context, owner, reponame, token string, aks models.AKS) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	templog := workflow.GetLogger(ctx)

	ctx = workflow.WithActivityOptions(ctx, ao)

	var a *activities.AKSImpl
	var ticketNumber string
	err := workflow.ExecuteActivity(ctx, a.GitCreateTicketAppDeployActivity, owner, reponame, token, aks.AKSClusterName, aks.SubscriptionID, aks.Location, aks.AKSAppName).Get(ctx, &ticketNumber)
	if err != nil {
		return "", err
	}

	templog.Info(utils.GitWorkflow, "Git Ticket Value is ", ticketNumber)

	err = workflow.ExecuteActivity(ctx, a.GitPollTicketActivity, ticketNumber, token).Get(ctx, &ticketNumber)
	if err != nil {
		return "", err
	}

	return ticketNumber, nil
}
