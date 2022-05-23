package ports

import (
	"context"
	"fmt"
	"net/http"
	"temporal-eg/common/constants"
	purchaseorch "temporal-eg/pkg/temporal/purchase-orch"

	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
)

func RunHttpServer(serverPort int, ctrl *Controller) error {
	router := gin.Default()

	router.POST("/order", ctrl.HandlePostOrder)

	return router.Run(fmt.Sprintf(":%d", serverPort))
}

type Controller struct {
	temporalClient client.Client
}

func NewController(tc client.Client) *Controller {
	return &Controller{
		temporalClient: tc,
	}
}

func (c *Controller) HandlePostOrder(ctx *gin.Context) {
	var createOrderRequest CreateOrderRequest
	if err := ctx.ShouldBindJSON(&createOrderRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: string(constants.PurchaseTaskQueueName),
	}

	workflowParams := purchaseorch.PurchaseWorkflowParams{
		ProductID: createOrderRequest.ProductID,
		UserID:    createOrderRequest.UserID,
		Quantity:  createOrderRequest.Quantity,
	}
	workflowRun, err := c.temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, purchaseorch.PurchaseWorkflow, workflowParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("error running workflow: %w", err)})
		return
	}

	var result purchaseorch.PurchaseWorkflowResponse
	err = workflowRun.Get(context.Background(), &result)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("error getting workflow result: %w", err)})
		return
	}

	createOrderResponse := CreateOrderResponse{
		OrderID: result.OrderID,
	}

	ctx.JSON(http.StatusOK, createOrderResponse)
}
