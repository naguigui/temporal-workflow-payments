package main

import (
	"fmt"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporal-eg/common/constants"
	"temporal-eg/pkg/temporal/sales"
	inmemory "temporal-eg/sales-service/adapter/in-memory"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		fmt.Println("failed to create temporal client: %w", err)
		os.Exit(-1)
	}
	defer c.Close()
	w := worker.New(c, string(constants.SalesTaskQueueName), worker.Options{})

	salesRepo := inmemory.NewRepository()
	salesActivity := sales.Activity{
		Repo: salesRepo,
	}

	w.RegisterActivity(salesActivity.ReserveProduct)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		fmt.Println("failed to run worker: %w", err)
		os.Exit(-1)
	}
}
