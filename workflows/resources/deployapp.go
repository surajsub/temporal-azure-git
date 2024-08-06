package resources

import (
	"github.com/surajsub/temporal-azure-git/activities"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/workflow"
	"time"
)

func DeployAppWorkflow(ctx workflow.Context, ipaddress string, aks models.AKS) (map[string]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 30,
	}
	templog := workflow.GetLogger(ctx)

	ctx = workflow.WithActivityOptions(ctx, ao)

	var a *activities.AKSImpl

	var kubeconfigpath string
	err := workflow.ExecuteActivity(ctx, a.KubeConfigActivity, aks).Get(ctx, &kubeconfigpath)
	if err != nil {
		return nil, err
	}

	var deployOutput map[string]string
	err = workflow.ExecuteActivity(ctx, a.DeployResourcesWithKubectlActivity, kubeconfigpath, "/tmp/", aks).Get(ctx, &deployOutput)
	if err != nil {
		return nil, err
	}

	templog.Info(utils.DeployAppWorkflow, "Deployment is complete  is ", deployOutput[utils.AKS_ID])
	return deployOutput, nil
}
