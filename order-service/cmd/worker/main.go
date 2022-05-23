package main

import (
	"context"
	"fmt"
	"os"

	"temporal-eg/common/constants"

	purchaseorch "temporal-eg/pkg/temporal/purchase-orch"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func main() {
	temporalClient, err := client.NewClient(client.Options{})
	if err != nil {
		fmt.Println("failed to initialize temporal client", err)
		os.Exit(-1)
	}
	defer temporalClient.Close()

	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: string(constants.PurchaseTaskQueueName),
	}

	workflowParams := purchaseorch.PurchaseWorkflowParams{
		ProductID: uuid.NewString(),
		UserID:    uuid.NewString(),
		Quantity:  3,
	}
	workflowRun, err := temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, purchaseorch.PurchaseWorkflow, workflowParams)
	if err != nil {
		fmt.Println("Error running workflow", err)
	}

	var result purchaseorch.PurchaseWorkflowResponse
	err = workflowRun.Get(context.Background(), &result)
	if err != nil {
		fmt.Println("Error getting workflow result", err)
	}
}
