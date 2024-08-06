package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/activity"
)

var provisioner utils.Provisioner

// This is the common vnet provisioner

func (a *AKSImpl) VNetInitActivity(ctx context.Context) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Init(dir + utils.VNET_DIR)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.INIT_ACTIVITY, engine)
	return output, nil
}

func (a *AKSImpl) VNetApplyActivity(ctx context.Context, vnetblock string, rg, location, env, vnetName string) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Apply(dir+utils.VNET_DIR,
		"-var", fmt.Sprintf("vnet_block=%s", vnetblock),
		"-var", fmt.Sprintf("vnet_name=%s", vnetName),
		"-var", fmt.Sprintf("region=%s", location),
		"-var", fmt.Sprintf("env=%s", env),
		"-var", fmt.Sprintf("rg_name=%s", rg))
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.APPLY_ACTIVITY, engine)
	return output, nil
}

func (a *AKSImpl) VNetOutputActivity(ctx context.Context) (map[string]string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	outputValues, err := provisioner.Output(dir + utils.VNET_DIR)
	if err != nil {
		return nil, err
	}
	activity.GetLogger(ctx).Info(utils.OUTPUT_ACTIVITY, engine)

	var tfOutput map[string]models.VNetCommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %w", err)
	}

	vnetOutput := map[string]string{
		"vnet_id":      tfOutput[utils.VNET_ID].Value,
		"vnet_guid_id": tfOutput[utils.VNET_GUID_ID].Value,
		"vnet_name":    tfOutput[utils.VNET_NAME].Value,
	}

	return vnetOutput, nil
}
