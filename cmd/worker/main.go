package worker

import (
	"log"

	"github.com/macneib/cedar-policy-manager/pkg/activities"
	"github.com/macneib/cedar-policy-manager/pkg/db"
	"github.com/macneib/cedar-policy-manager/pkg/workflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func StartWorker() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	w := worker.New(c, "cedar_policy_task_queue", worker.Options{})

	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	activity := &activities.CedarPolicyActivities{DB: dbConn}
	w.RegisterWorkflow(workflows.ManageCedarPolicyWorkflow)
	w.RegisterActivity(activity.CreatePolicyActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatal(err)
	}
}
