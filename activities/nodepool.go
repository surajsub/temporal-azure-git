package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/activity"
	"log"
)

// This is the common vnet provisioner
func (a *AKSImpl) NodePoolInitActivity(ctx context.Context) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Init(dir + utils.NP_DIR)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.INIT_ACTIVITY, engine)
	return output, nil
}

func (a *AKSImpl) NodePoolApplyActivity(ctx context.Context, aksid, akssubnetid string, aks models.AKS) (string, error) {

	activity.GetLogger(ctx).Info("in the nodepool activity with value  ", aksid)

	log.Printf("the value of the id of the aks cluster is %s\n", aksid)
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Apply(dir+utils.NP_DIR,
		"-var", fmt.Sprintf("aks_id=%s", aksid),
		"-var", fmt.Sprintf("env=%s", aks.Env),
		"-var", fmt.Sprintf("aks_version=%s", aks.AKSVersion),
		"-var", fmt.Sprintf("vm_size=%s", aks.AKSVmSize),
		"-var", fmt.Sprintf("aks_subnet_id=%s", akssubnetid))
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.APPLY_ACTIVITY, engine)
	return output, nil
}

func (a *AKSImpl) NodePoolOutputActivity(ctx context.Context) (map[string]string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	outputValues, err := provisioner.Output(dir + utils.NP_DIR)
	if err != nil {
		return nil, err
	}
	activity.GetLogger(ctx).Info(utils.OUTPUT_ACTIVITY, engine)

	var tfOutput map[string]models.NPCommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %w", err)
	}

	npOutput := map[string]string{
		"np_id": tfOutput[utils.RG_ID].Value,
	}

	return npOutput, nil
}
