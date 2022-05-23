package main

import (
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporal-eg/common/constants"
	inmemory "temporal-eg/payment-service/adapter/in-memory"
	"temporal-eg/pkg/log"
	"temporal-eg/pkg/temporal/payment"
)

func main() {
	logger := log.New()
	c, err := client.NewClient(client.Options{})
	if err != nil {
		logger.Errorf("failed to create temporal client: %w", err)
		os.Exit(-1)
	}
	defer c.Close()
	w := worker.New(c, string(constants.PaymentTaskQueueName), worker.Options{})

	paymentRepo := inmemory.NewRepository()
	paymentActivity := payment.Activity{
		Repo: paymentRepo,
	}

	w.RegisterActivity(paymentActivity.CreatePaymentIntent)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		logger.Errorf("failed to run worker: %w", err)
		os.Exit(-1)
	}
}
