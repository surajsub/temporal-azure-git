package resources

import (
	"github.com/surajsub/temporal-azure-git/activities"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/workflow"
	"time"
)

func RGWorkflow(ctx workflow.Context, aks models.AKS) (map[string]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	templog := workflow.GetLogger(ctx)

	var a *activities.AKSImpl

	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, a.RGInitActivity).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, a.RGApplyActivity, aks.Location, aks.ResourceGroup, aks.Env).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	var rgOutput map[string]string
	err = workflow.ExecuteActivity(ctx, a.RGOutputActivity).Get(ctx, &rgOutput)
	if err != nil {
		return nil, err
	}

	templog.Info(utils.RGWorkflow, "RG Value is ", rgOutput[utils.RG_ID])
	return rgOutput, nil
}
