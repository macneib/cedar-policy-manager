package worker

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func InitDB() {
	var err error
	db, err = sqlx.Connect("postgres", "user=youruser dbname=yourdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func CreatePolicyActivity(ctx context.Context, tenantID string, policyName string, policyDocument string) (string, error) {
	var policyID int
	query := `INSERT INTO policies (tenant_id, policy_name, policy_document) VALUES ($1, $2, $3) RETURNING policy_id`
	err := db.QueryRowContext(ctx, query, tenantID, policyName, policyDocument).Scan(&policyID)
	if err != nil {
		return "", err
	}
	return "Policy Created", nil
}
