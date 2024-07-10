package worker

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func ManageCedarPolicyWorkflow(ctx workflow.Context, tenantID string, policyName string, policyDocument string) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	err := workflow.ExecuteActivity(ctx, CreatePolicyActivity, tenantID, policyName, policyDocument).Get(ctx, &result)
	if err != nil {
		return err
	}

	return nil
}
