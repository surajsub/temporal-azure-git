package resources

import (
	"github.com/surajsub/temporal-azure-git/activities"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/workflow"
	"time"
)

func NodePoolWorkflow(ctx workflow.Context, aksid, akssubnetid string, aks models.AKS) (map[string]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 30,
	}
	templog := workflow.GetLogger(ctx)

	ctx = workflow.WithActivityOptions(ctx, ao)

	var a *activities.AKSImpl

	workflow.GetLogger(ctx).Debug("The Id of the Kubernetes cluster is %s and the value of the subnet is %s", aksid, akssubnetid)

	err := workflow.ExecuteActivity(ctx, a.NodePoolInitActivity).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, a.NodePoolApplyActivity, aksid, akssubnetid, aks).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	var npOutput map[string]string
	err = workflow.ExecuteActivity(ctx, a.NodePoolOutputActivity).Get(ctx, &npOutput)
	if err != nil {
		return nil, err
	}

	templog.Info(utils.NodePoolWorkflow, "VNET Value is ", npOutput[utils.AKS_ID])
	return npOutput, nil
}
