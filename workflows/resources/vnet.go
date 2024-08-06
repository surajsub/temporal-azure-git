package resources

import (
	"github.com/surajsub/temporal-azure-git/activities"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/workflow"
	"time"
)

func VNETWorkflow(ctx workflow.Context, vnetBlock string, aks models.AKS) (map[string]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	templog := workflow.GetLogger(ctx)
	var a *activities.AKSImpl

	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, a.VNetInitActivity).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, a.VNetApplyActivity, vnetBlock, aks.ResourceGroup, aks.Location, aks.Env, aks.VnetName).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	var vnetOutput map[string]string
	err = workflow.ExecuteActivity(ctx, a.VNetOutputActivity).Get(ctx, &vnetOutput)
	if err != nil {
		return nil, err
	}

	templog.Info(utils.VNetWorkflow, "VNET Value is ", vnetOutput[utils.VNET_ID])
	return vnetOutput, nil
}
