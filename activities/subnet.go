package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/activity"
)

func (a *AKSImpl) SubnetInitActivity(ctx context.Context) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Init(dir + utils.SUBNET_DIR)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.INIT_ACTIVITY, engine)
	return output, nil
}

func (a *AKSImpl) SubnetApplyActivity(ctx context.Context, vnet, rgname string) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Apply(dir+utils.SUBNET_DIR,
		"-var", fmt.Sprintf("vnet_name=%s", vnet),
		"-var", fmt.Sprintf("rg_name=%s", rgname))
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.APPLY_ACTIVITY, engine)
	return output, nil
}

func (a *AKSImpl) SubnetOutputActivity(ctx context.Context) (map[string]string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	outputValues, err := provisioner.Output(dir + utils.SUBNET_DIR)
	if err != nil {
		return nil, err
	}
	activity.GetLogger(ctx).Info(utils.OUTPUT_ACTIVITY, engine)

	var tfOutput map[string]models.SubnetCommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %w", err)
	}

	subnetOutput := map[string]string{
		"aks_subnet_id":     tfOutput["aks_subnet_id"].Value,
		"aks_app_subnet_id": tfOutput["aks_app_subnet_id"].Value,
		"aks_subnet_name":   tfOutput["aks_subnet_net"].Value,
	}

	return subnetOutput, nil
}
