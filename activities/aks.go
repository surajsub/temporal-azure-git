package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/activity"
)

// This is the common vnet provisioner

func (a *AKSImpl) AKSInitActivity(ctx context.Context) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Init(dir + utils.AKS_DIR)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.INIT_ACTIVITY, engine)
	return output, nil
}

// Azure Kubernetes Service
func (a *AKSImpl) AKSApplyActivity(ctx context.Context, rgID, subnetID, userId string, aks models.AKS) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Apply(dir+utils.AKS_DIR,
		"-var", fmt.Sprintf("location=%s", aks.Location),
		"-var", fmt.Sprintf("rg_name=%s", aks.ResourceGroup),
		"-var", fmt.Sprintf("aks_name=%s", aks.AKSClusterName),
		"-var", fmt.Sprintf("rg_id=%s", rgID),
		"-var", fmt.Sprintf("aks_subnet_id=%s", subnetID),
		"-var", fmt.Sprintf("env=%s", aks.Env),
		"-var", fmt.Sprintf("user_aid_id=%s", userId))
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.APPLY_ACTIVITY, engine)
	return output, nil
}

func (a *AKSImpl) AKSOutputActivity(ctx context.Context) (map[string]string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	outputValues, err := provisioner.Output(dir + utils.AKS_DIR)
	if err != nil {
		return nil, err
	}
	activity.GetLogger(ctx).Info(utils.OUTPUT_ACTIVITY, engine)

	var tfOutput map[string]models.AKSCommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %w", err)
	}

	aksOutput := map[string]string{
		"aks_id": tfOutput[utils.AKS_ID].Value,
	}

	return aksOutput, nil
}
