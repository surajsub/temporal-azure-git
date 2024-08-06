package activities

import (
	"context"
	"github.com/surajsub/temporal-azure-git/models"
	"go.temporal.io/sdk/activity"
	"log"
	"os/exec"
)

func (a *AKSImpl) KubeConfigActivity(ctx context.Context, aks models.AKS) (string, error) {

	activity.GetLogger(ctx).Info("Starting the KubeConfig work")
	kubeconfigCmd := exec.Command("az", "aks", "get-credentials", "--resource-group", aks.ResourceGroup, "--name", aks.AKSClusterName, "--file", "/tmp/kubeconfig") //nolint:gosec
	err := kubeconfigCmd.Run()
	if err != nil {
		return "", err
	}
	return "/tmp/kubeconfig", nil

}

func (a *AKSImpl) DeployResourcesWithKubectlActivity(ctx context.Context, kubeconfigPath, yamlFilePath string, aks models.AKS) error {

	yamlFilePath = "/tmp/aks-store-quickstart.yaml"
	activity.GetLogger(ctx).Info("Starting the Application Deploy Work with the following values", kubeconfigPath, "and the yamlfile path is", kubeconfigPath, yamlFilePath)
	log.Println("calling the following command ")
	log.Printf("kubectl --kubeconfig %s apply -f %s\n", kubeconfigPath, yamlFilePath)
	kubectlCmd := exec.Command("kubectl", "--kubeconfig", kubeconfigPath, "apply", "-f", "/tmp/aks-store-quickstart.yaml")
	err := kubectlCmd.Run()
	if err != nil {
		return err
	}
	return nil
}
