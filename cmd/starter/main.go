package main

import (
	"context"
	"log"

	"cedar-policy-manager/worker" // Import the worker package

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "cedar_policy_workflow",
		TaskQueue: "cedar_policy_task_queue",
	}

	tenantID := "tenant1-uuid"
	policyName := "AllowS3Access"
	policyDocument := `{"Version": "2024-07-10", "Statement": [{"Effect": "Allow", "Action": "s3:*", "Resource": "*"}]}`

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, worker.ManageCedarPolicyWorkflow, tenantID, policyName, policyDocument)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
