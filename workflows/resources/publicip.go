package resources

import (
	"github.com/surajsub/temporal-azure-git/activities"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/workflow"
	"time"
)

func PublicIPWorkflow(ctx workflow.Context, aks models.AKS) (map[string]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	templog := workflow.GetLogger(ctx)

	var a *activities.AKSImpl

	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, a.PublicIPInitActivity).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, a.PublicIPApplyActivity, aks).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	var pioutput map[string]string
	err = workflow.ExecuteActivity(ctx, a.PublicIPOutputActivity).Get(ctx, &pioutput)
	if err != nil {
		return nil, err
	}

	templog.Info(utils.VNetWorkflow, "Public IP  Value is ", pioutput[utils.PIP_IP])
	return pioutput, nil
}
