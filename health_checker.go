package sqs

import (
	"context"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type HealthChecker struct {
	Client    *sqs.SQS
	QueueName *string
	Service   string
}

func NewHealthChecker(client *sqs.SQS, queueName string) *HealthChecker {
	return NewSQSHealthChecker(client, queueName, "sqs")
}
func NewSQSHealthChecker(client *sqs.SQS, name string, queueName string) *HealthChecker {
	return &HealthChecker{client, &queueName, name}
}

func (h *HealthChecker) Name() string {
	return h.Service
}

func (h *HealthChecker) Check(ctx context.Context) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	_, err := h.Client.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: h.QueueName,
	})
	return res, err
}

func (h *HealthChecker) Build(ctx context.Context, data map[string]interface{}, err error) map[string]interface{} {
	if err == nil {
		return data
	}
	data["error"] = err.Error()
	return data
}
