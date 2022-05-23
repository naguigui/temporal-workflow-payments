package main

import (
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporal-eg/common/constants"
	inmemory "temporal-eg/order-service/adapter/in-memory"
	purchaseorch "temporal-eg/pkg/temporal/purchase-orch"

	"temporal-eg/pkg/log"
)

func main() {
	logger := log.New()
	c, err := client.NewClient(client.Options{})
	if err != nil {
		logger.Errorf("failed to create temporal client: %w", err)
		os.Exit(-1)
	}
	defer c.Close()
	w := worker.New(c, string(constants.PurchaseTaskQueueName), worker.Options{})

	purchaseOrchActivity := purchaseorch.Activity{
		Repo: inmemory.NewRepository(),
	}
	w.RegisterActivity(purchaseOrchActivity.CreateOrder)
	w.RegisterWorkflow(purchaseorch.PurchaseWorkflow)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		logger.Errorf("failed to run order worker: %w", err)
		os.Exit(-1)
	}
}
