package handlers

import (
	"context"
	"go.temporal.io/sdk/client"
	"log"
)

func StartWorkflow() {
	// Create Temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Define workflow options
	workflowOptions := client.StartWorkflowOptions{
		ID:        "terraform-workflow",   // Name of the workflow that will be visible in the Temporal UI
		TaskQueue: "terraform-task-queue", // Queue Name - This can be made dynamic
	}

	// Start the workflow
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, TerraformWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Printf("Started workflow with ID: %s and Run ID: %s", we.GetID(), we.GetRunID())

	var result interface{}
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatal("Unable to get workflow result", err)
	}
	log.Printf("Workflow result: %v", result)
}
