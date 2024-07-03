package worker

import (
	"github.com/surajsub/temporal-terraform-example/activities"
	"github.com/surajsub/temporal-terraform-example/handlers"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func RunWorker() {
	// Create Temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Create a new worker
	w := worker.New(c, "terraform-task-queue", worker.Options{})

	// Register the workflow and activities
	w.RegisterWorkflow(handlers.TerraformWorkflow)
	w.RegisterActivity(activities.TerraformInitActivity)
	w.RegisterActivity(activities.TerraformApplyActivity)
	w.RegisterActivity(activities.TerraformOutputActivity)

	// Start the worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
