package workflows

import (
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"github.com/surajsub/temporal-azure-git/workflows/resources"
	"go.temporal.io/sdk/workflow"
	"log"
	"time"
)

func ParentWorkflow(ctx workflow.Context, vpc string, aks models.AKS, git models.GitData) (map[string]interface{}, error) {
	cwo := workflow.ChildWorkflowOptions{
		WorkflowExecutionTimeout: time.Hour,
		WorkflowRunTimeout:       time.Hour * 36,
	}
	ctx = workflow.WithChildOptions(ctx, cwo)
	workflowID := workflow.GetInfo(ctx).OriginalRunID

	log.Printf("Printing the wortkflow id from the PARENT %s\n", workflowID)

	// Start RG Workflow
	var rgOutput map[string]string
	err := workflow.ExecuteChildWorkflow(ctx, resources.RGWorkflow, aks).Get(ctx, &rgOutput)
	if err != nil {
		return nil, err
	}
	workflow.GetLogger(ctx).Info("RG created", "rg_id", rgOutput[utils.RG_ID])

	// Start VNET Workflow
	var vnetOutput map[string]string
	err = workflow.ExecuteChildWorkflow(ctx, resources.VNETWorkflow, vpc, aks).Get(ctx, &vnetOutput)
	if err != nil {
		return nil, err
	}
	workflow.GetLogger(ctx).Info("VNET created", "VNet ID", vnetOutput[utils.VNET_ID])
	workflow.GetLogger(ctx).Info("VNET Created", "VNET GUID ID ", vnetOutput[utils.VNET_GUID_ID])

	// Start subnet flow
	var subnetOutput map[string]string
	err = workflow.ExecuteChildWorkflow(ctx, resources.SubnetWorkflow, vnetOutput[utils.VNET_NAME], aks).Get(ctx, &subnetOutput)
	if err != nil {
		return nil, err
	}
	workflow.GetLogger(ctx).Info("Subnet Created", "Subnet  ID ", subnetOutput[utils.SUBNET_1])

	// Start the Managed Identity
	var mioutput map[string]string
	err = workflow.ExecuteChildWorkflow(ctx, resources.ManagedIdentityWorkflow, rgOutput[utils.RG_ID], aks).Get(ctx, &mioutput)
	if err != nil {
		log.Println("Failed to execute the ManagedIdentity Workflow")
		return nil, err
	}

	// Start GIT Flow - When this workflow starts it creates a ticket that must be approved by a human.
	// The only action needed is to close the ticket.
	// TODO - Add Logic to make sure that the ticket is approved.
	var ticket string
	err = workflow.ExecuteChildWorkflow(ctx, resources.GitApprovalWorkflow, git.Owner, git.RepoName, git.GitToken, aks).Get(ctx, &ticket)
	if err != nil {
		return nil, err
	}
	log.Printf("the ticket that was closed is %v \n", &ticket)

	// Start the AKS workflow
	var aksOutput map[string]string
	err = workflow.ExecuteChildWorkflow(ctx, resources.AKSWorkflow, rgOutput[utils.RG_ID], subnetOutput[utils.AKS_SUBNET_ID], mioutput[utils.MI_ID], aks).Get(ctx, &aksOutput)
	if err != nil {
		log.Println("Failed to execute the AKS Workflow")
		return nil, err
	}

	// Start the nodepool workflow

	var npOutput map[string]string
	err = workflow.ExecuteChildWorkflow(ctx, resources.NodePoolWorkflow, aksOutput[utils.AKS_ID], subnetOutput[utils.AKS_SUBNET_ID], aks).Get(ctx, &npOutput)
	if err != nil {
		log.Println("Failed to execute the AKS Workflow")
		return nil, err
	}

	// Start the public ip workflow

	var pipOutput map[string]string
	err = workflow.ExecuteChildWorkflow(ctx, resources.PublicIPWorkflow, aks).Get(ctx, &pipOutput)
	if err != nil {
		log.Println("Failed to execute the Public IP Workflow")
		return nil, err
	}

	// Start the DNS Flow

	var dnsOutput map[string]string
	err = workflow.ExecuteChildWorkflow(ctx, resources.DNSWorkflow, pipOutput[utils.PIP_IP], aks).Get(ctx, &dnsOutput)
	if err != nil {
		log.Println("Failed to execute the DNS Registration  Workflow")
		return nil, err
	}

	// Start the GIT App Deploy Approval Workflow

	var appticket string
	err = workflow.ExecuteChildWorkflow(ctx, resources.GitAppDeployWorkflow, git.Owner, git.RepoName, git.GitToken, aks).Get(ctx, &appticket)
	if err != nil {
		return nil, err
	}
	log.Printf("the ticket that was closed is %v \n", &ticket)

	var deployOutput map[string]string
	err = workflow.ExecuteChildWorkflow(ctx, resources.DeployAppWorkflow, pipOutput[utils.PIP_IP], aks).Get(ctx, &deployOutput)
	if err != nil {
		log.Println("Failed to execute the AKS Workflow")
		return nil, err
	}

	/*


		// Start the kubectl generation flow

			// Start the Managed Identity

			var mioutput map[string]string
			err = workflow.ExecuteChildWorkflow(ctx, resources.ManagedIdentityWorkflow, rgOutput[utils.RG_ID]).Get(ctx, &mioutput)
			if err != nil {
				log.Println("Failed to execute the ManagedIdentity Workflow")
				return nil, err
			}


	*/
	results := map[string]interface{}{
		"AzWorkflow": rgOutput,
	}

	// Start the AKS

	// Start the Node Pool

	return results, nil
}
