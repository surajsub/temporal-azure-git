package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/activity"
)

// This is the common DNS provisioner

func (a *AKSImpl) DNSInitActivity(ctx context.Context) (string, error) {

	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Init(dir + utils.DNS_DIR)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.INIT_ACTIVITY, engine)
	return output, nil
}

func (a *AKSImpl) DNSApplyActivity(ctx context.Context, ipaddress string, req models.AKS) (string, error) {

	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	output, err := provisioner.Apply(dir+utils.DNS_DIR,
		"-var", fmt.Sprintf("rg_name=%s", req.ResourceGroup),
		"-var", fmt.Sprintf("dns_name=%s", req.AKSDnsName),
		"-var", fmt.Sprintf("temporal_pip=%s", ipaddress))
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info(utils.APPLY_ACTIVITY, engine)
	return output, nil
}

func (a *AKSImpl) DNSOutputActivity(ctx context.Context) (map[string]string, error) {
	provisioner, engine, dir := utils.GetProvisioner(provisioner)
	outputValues, err := provisioner.Output(dir + utils.DNS_DIR)
	if err != nil {
		return nil, err
	}
	activity.GetLogger(ctx).Info(utils.OUTPUT_ACTIVITY, engine)

	var tfOutput map[string]models.DNSCommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %w", err)
	}

	pipOutput := map[string]string{
		"dns_ip":     tfOutput["dns_id"].Value,
		"dns_a_id":   tfOutput["dns_a_id"].Value,
		"dns_a_fqdn": tfOutput["dns_a_fqdn"].Value,
	}

	return pipOutput, nil
}
