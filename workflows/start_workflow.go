package workflows

import (
	"context"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/utils"
	"go.temporal.io/sdk/client"
	"log"
	"time"
)

func StartWorkflow(vpc string, aks models.AKS, git models.GitData) {
	// Create Temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Define workflow options
	workflowOptions := client.StartWorkflowOptions{
		ID:                       "parent-workflow-azure",   // Name of the workflow that will be visible in the Temporal UI
		TaskQueue:                utils.WORKFLOW_TASK_QUEUE, // Queue Name - This can be made dynamic
		WorkflowExecutionTimeout: 60 * time.Minute,
	}

	// Start the workflow
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, ParentWorkflow, vpc, aks, git)
	if err != nil {
		log.Panicln("Unable to execute workflow", err)
	}

	log.Printf("Started workflow with ID: %s and Run ID: %s", we.GetID(), we.GetRunID())

	var result interface{}
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Panic("Unable to get workflow result", err)
	}
	log.Printf("Workflow result: %v", result)
}
