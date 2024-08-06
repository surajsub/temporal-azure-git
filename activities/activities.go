package activities

import "context"
import "github.com/surajsub/temporal-azure-git/models"

type AKSImpl struct{}

type AKSActivities interface {
	RGInitActivity(ctx context.Context) (string, error)
	RGApplyActivity(ctx context.Context, request models.AKS) (string, error)
	RGOutputActivity(ctx context.Context) (map[string]string, error)

	VnetInitActivity(ctx context.Context) (string, error)
	VnetApplyActivity(ctx context.Context, request models.AKS) (string, error)
	VnetOutputActivity(ctx context.Context) (map[string]string, error)

	SubnetInitActivity(ctx context.Context) (string, error)
	SubnetApplyActivity(ctx context.Context, request models.AKS) (string, error)
	SubnetOutputActivity(ctx context.Context) (map[string]string, error)

	MIInitActivity(ctx context.Context) (string, error)
	MIApplyActivity(ctx context.Context, request models.AKS) (string, error)
	MIOutputActivity(ctx context.Context) (map[string]string, error)

	GitCreateTicketActivity(ctx context.Context) (string, error)
	GitPollTicketActivity(ctx context.Context) (string, error)
	GitCreateTicketAppDeployActivity(ctx context.Context, owner, reponame, token, clustername, subid, location string) (string, error)

	AKSInitActivity(ctx context.Context) (string, error)
	AKSApplyActivity(ctx context.Context, request models.AKS, rgId, subnetId, userId string) (string, error)
	AKSOutputActivity(ctx context.Context) (map[string]string, error)

	NodePoolInitActivity(ctx context.Context) (string, error)
	NodePoolApplyActivity(ctx context.Context, request models.AKS) (string, error)
	NodePoolOutputActivity(ctx context.Context) (map[string]string, error)

	PublicIPInitActivity(ctx context.Context) (string, error)
	PublicIPApplyActivity(ctx context.Context, request models.AKS) (string, error)
	PublicIPOutputActivity(ctx context.Context) (map[string]string, error)

	DNSInitActivity(ctx context.Context) (string, error)
	DNSApplyActivity(ctx context.Context, request models.AKS, ipaddress string) (string, error)
	DNSOutputActivity(ctx context.Context) (map[string]string, error)

	// generate the kubeconfig file and then deploy the application

	KubeConfigActivity(ctx context.Context, aks models.AKS) (string, error)
	DeployResourcesWithKubectlActivity(ctx context.Context, filename, filepath string) error
}
