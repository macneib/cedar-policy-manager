package activities

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type CedarPolicyActivities struct {
	DB *sqlx.DB
}

func (a *CedarPolicyActivities) CreatePolicyActivity(ctx context.Context, tenantID string, policyName string, policyDocument string) (string, error) {
	var policyID int
	query := `INSERT INTO policies (tenant_id, policy_name, policy_document) VALUES ($1, $2, $3) RETURNING policy_id`
	err := a.DB.QueryRowContext(ctx, query, tenantID, policyName, policyDocument).Scan(&policyID)
	if err != nil {
		return "", err
	}
	return "Policy Created", nil
}
