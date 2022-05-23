package main

import (
	"fmt"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporal-eg/common/constants"
	inmemory "temporal-eg/order-service/adapter/in-memory"
	purchaseorch "temporal-eg/pkg/temporal/purchase-orch"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		fmt.Println("failed to create temporal client: %w", err)
		os.Exit(-1)
	}
	defer c.Close()

	purchaseActivity := purchaseorch.Activity{
		Repo: inmemory.NewRepository(),
	}

	w := worker.New(c, string(constants.PurchaseTaskQueueName), worker.Options{})
	w.RegisterWorkflow(purchaseorch.PurchaseWorkflow)
	w.RegisterActivity(purchaseActivity.CreateOrder)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		fmt.Println("failed to run worker: %w", err)
		os.Exit(-1)
	}
}
