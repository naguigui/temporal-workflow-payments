package main

import (
	"os"
	"temporal-eg/order-service/ports"
	"temporal-eg/pkg/log"

	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

func main() {
	logger := log.New()
	temporalClient, err := client.NewClient(client.Options{})
	if err != nil {
		logger.Errorf("failed initialize temporal client", zap.Error(err))
		os.Exit(-1)
	}
	defer temporalClient.Close()

	ctrl := ports.NewController(temporalClient)
	if err := ports.RunHttpServer(8080, ctrl); err != nil {
		logger.Errorf("failed starting http server", zap.Error(err))
		os.Exit(1)
	}
}
