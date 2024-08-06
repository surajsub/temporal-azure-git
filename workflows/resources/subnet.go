package resources

import (
	"github.com/surajsub/temporal-azure-git/activities"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/workflow"
	"time"
)

func SubnetWorkflow(ctx workflow.Context, vnetName string, aks models.AKS) (map[string]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	templog := workflow.GetLogger(ctx)

	var a *activities.AKSImpl

	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, a.SubnetInitActivity).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, a.SubnetApplyActivity, vnetName, aks.ResourceGroup).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	var subnetOutput map[string]string
	err = workflow.ExecuteActivity(ctx, a.SubnetOutputActivity).Get(ctx, &subnetOutput)
	if err != nil {
		return nil, err
	}

	templog.Info(utils.VNetWorkflow, "VNET Value is ", subnetOutput[utils.SUBNETID])
	return subnetOutput, nil
}
