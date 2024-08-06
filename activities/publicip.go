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

func (a *AKSImpl) PublicIPInitActivity(ctx context.Context) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Init(dir + utils.PIP_DIR)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.INIT_ACTIVITY, engine)
	return output, nil
}

func (a *AKSImpl) PublicIPApplyActivity(ctx context.Context, aks models.AKS) (string, error) {

	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Apply(dir+utils.PIP_DIR,
		"-var", fmt.Sprintf("rg_location=%s", aks.Location),
		"-var", fmt.Sprintf("env=%s", aks.Env),
		"-var", fmt.Sprintf("rg_name=%s", aks.ResourceGroup))
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.APPLY_ACTIVITY, engine)
	return output, nil
}

func (a *AKSImpl) PublicIPOutputActivity(ctx context.Context) (map[string]string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	outputValues, err := provisioner.Output(dir + utils.PIP_DIR)
	if err != nil {
		return nil, err
	}
	activity.GetLogger(ctx).Info(utils.OUTPUT_ACTIVITY, engine)

	var tfOutput map[string]models.PIPCommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %w", err)
	}

	pipOutput := map[string]string{
		"public_ip":    tfOutput["public_ip"].Value,
		"public_ip_id": tfOutput[utils.PIP_ID].Value,
	}

	return pipOutput, nil
}
