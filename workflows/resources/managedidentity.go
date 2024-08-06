package resources

import (
	"github.com/surajsub/temporal-azure-git/activities"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/workflow"
	"time"
)

func ManagedIdentityWorkflow(ctx workflow.Context, rgID string, aks models.AKS) (map[string]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	templog := workflow.GetLogger(ctx)

	var a *activities.AKSImpl

	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, a.ManagedIdentityInitActivity).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, a.ManagedIdentityApplyActivity, rgID, aks).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	var miOutput map[string]string
	err = workflow.ExecuteActivity(ctx, a.ManagedIdentityOutputActivity).Get(ctx, &miOutput)
	if err != nil {
		return nil, err
	}

	templog.Info(utils.VNetWorkflow, "VNET Value is ", miOutput[utils.MI_ID])
	return miOutput, nil
}
