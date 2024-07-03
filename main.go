package main

import (
	"github.com/surajsub/temporal-terraform-example/handlers"
	"github.com/surajsub/temporal-terraform-example/worker"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("expected 'worker' or 'starter' subcommands")
	}

	switch os.Args[1] {
	case "worker":
		worker.RunWorker()
	case "starter":
		handlers.StartWorkflow()
	default:
		log.Fatal("expected 'worker' or 'starter' subcommands")
	}
}
