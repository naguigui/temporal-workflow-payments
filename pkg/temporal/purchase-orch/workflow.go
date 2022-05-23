package purchaseorch

import (
	"fmt"
	"temporal-eg/common/constants"
	"temporal-eg/pkg/temporal/payment"
	"temporal-eg/pkg/temporal/sales"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type PurchaseWorkflowParams struct {
	ProductID string
	UserID    string
	Quantity  int
}

type PurchaseWorkflowResponse struct {
	OrderID string
}

func PurchaseWorkflow(ctx workflow.Context, p PurchaseWorkflowParams) (*PurchaseWorkflowResponse, error) {
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    500,
	}
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retrypolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var purchaseOrchActivity *Activity
	future := workflow.ExecuteActivity(ctx, purchaseOrchActivity.CreateOrder, p.ProductID, p.UserID, p.Quantity)
	var createOrderRes string
	if err := future.Get(ctx, &createOrderRes); err != nil {
		return nil, fmt.Errorf("failed executing create order activity")
	}

	var salesActivity *sales.Activity
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retrypolicy,
		TaskQueue:           string(constants.SalesTaskQueueName),
	})

	if err := workflow.ExecuteActivity(ctx, salesActivity.ReserveProduct, createOrderRes, p.Quantity).Get(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed executing reserve product activity: %w", err)
	}

	var paymentActivity *payment.Activity
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retrypolicy,
		TaskQueue:           string(constants.PaymentTaskQueueName),
	})
	createPaymentIntentParams := payment.CreatePaymentIntentParams{
		Amount:   float64(3 * p.Quantity),
		Currency: "USD",
	}
	var createPaymentResponse *payment.CreatePaymentIntentResponse
	if err := workflow.ExecuteActivity(ctx, paymentActivity.CreatePaymentIntent, createPaymentIntentParams).Get(ctx, &createPaymentResponse); err != nil {
		return nil, fmt.Errorf("failed executing payment intent: %w", err)
	}

	res := &PurchaseWorkflowResponse{
		OrderID: createOrderRes,
	}
	return res, nil
}
