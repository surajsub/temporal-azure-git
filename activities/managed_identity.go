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

func (a *AKSImpl) ManagedIdentityInitActivity(ctx context.Context) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Init(dir + utils.MI_DIR)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.INIT_ACTIVITY, engine)
	return output, nil
}

// Managed Identity

func (a *AKSImpl) ManagedIdentityApplyActivity(ctx context.Context, rgId string, aks models.AKS) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Apply(dir+utils.MI_DIR,
		"-var", fmt.Sprintf("location=%s", aks.Location),
		"-var", fmt.Sprintf("rgname=%s", aks.ResourceGroup),
		"-var", fmt.Sprintf("rg_id=%s", rgId))
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.APPLY_ACTIVITY, engine)
	return output, nil
}

func (a *AKSImpl) ManagedIdentityOutputActivity(ctx context.Context) (map[string]string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	outputValues, err := provisioner.Output(dir + utils.MI_DIR)
	if err != nil {
		return nil, err
	}
	activity.GetLogger(ctx).Info(utils.OUTPUT_ACTIVITY, engine)

	var tfOutput map[string]models.MICommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %w", err)
	}

	miOutput := map[string]string{
		"mi_id":           tfOutput[utils.MI_ID].Value,
		"mi_principal_id": tfOutput[utils.MI_PRINCIPAL_ID].Value,
		"mi_client_id":    tfOutput[utils.MI_CLIENT_ID].Value,
	}

	return miOutput, nil
}
