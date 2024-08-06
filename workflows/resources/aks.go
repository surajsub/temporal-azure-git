package resources

import (
	"github.com/surajsub/temporal-azure-git/activities"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/workflow"
	"time"
)

func AKSWorkflow(ctx workflow.Context, rgId, subnetId, userAidId string, aks models.AKS) (map[string]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 30,
	}
	templog := workflow.GetLogger(ctx)

	var a *activities.AKSImpl

	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, a.AKSInitActivity).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, a.AKSApplyActivity, rgId, subnetId, userAidId, aks).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	var aksOutput map[string]string
	err = workflow.ExecuteActivity(ctx, a.AKSOutputActivity).Get(ctx, &aksOutput)
	if err != nil {
		return nil, err
	}

	templog.Info(utils.AKSWorkflow, "VNET Value is ", aksOutput[utils.AKS_ID])
	return aksOutput, nil
}
