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

func (a *AKSImpl) RGInitActivity(ctx context.Context) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Init(dir + utils.RG_DIR)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.INIT_ACTIVITY, engine)
	return output, nil
}

func (a *AKSImpl) RGApplyActivity(ctx context.Context, location, name, env string) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Apply(dir+utils.RG_DIR,
		"-var", fmt.Sprintf("region=%s", location),
		"-var", fmt.Sprintf("name=%s", name),
		"-var", fmt.Sprintf("env=%s", env))
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.APPLY_ACTIVITY, engine)
	return output, nil
}

func (a *AKSImpl) RGOutputActivity(ctx context.Context) (map[string]string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	outputValues, err := provisioner.Output(dir + utils.RG_DIR)
	if err != nil {
		return nil, err
	}
	activity.GetLogger(ctx).Info(utils.OUTPUT_ACTIVITY, engine)

	var tfOutput map[string]models.RGCommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %w", err)
	}

	rgOutput := map[string]string{
		"rg_id":   tfOutput[utils.RG_ID].Value,
		"rg_name": tfOutput[utils.RG_NAME].Value,
	}

	return rgOutput, nil
}
