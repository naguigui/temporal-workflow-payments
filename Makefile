.PHONY: order-web
order-web:
	go run order-service/cmd/web/main.go

.PHONY: order-worker
order-worker:
	go run order-service/cmd/worker/main.go

.PHONY: sales-worker
sales-dev:
	go run sales-service/cmd/worker/main.go

.PHONY: payment-worker
payment-dev:
	go run payment-service/cmd/worker/main.go
