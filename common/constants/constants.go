package constants

type TaskQueueName string

const (
	PurchaseTaskQueueName TaskQueueName = "purchase"
	SalesTaskQueueName    TaskQueueName = "sales"
	PaymentTaskQueueName  TaskQueueName = "payment"
)
