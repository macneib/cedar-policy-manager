package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/macneib/cedar-policy-manager/worker" // Import the worker package
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	w := worker.New(c, "cedar_policy_task_queue", worker.Options{})

	w.RegisterWorkflow(worker.ManageCedarPolicyWorkflow)
	w.RegisterActivity(worker.CreatePolicyActivity)

	worker.InitDB()

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatal(err)
	}
}
