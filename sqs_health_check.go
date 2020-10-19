package sqs

import (
	"context"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSSubHealthCheck struct {
	Client    *sqs.SQS
	QueueName *string
	Service   string
}

func NewSQSHealthCheck(client *sqs.SQS, name string, queueName string) *SQSSubHealthCheck {
	return &SQSSubHealthCheck{client, &queueName, name}
}

func (h *SQSSubHealthCheck) Name() string {
	return h.Service
}

func (h *SQSSubHealthCheck) Check(ctx context.Context) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	_, err := h.Client.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: h.QueueName,
	})
	return res, err
}

func (h *SQSSubHealthCheck) Build(ctx context.Context, data map[string]interface{}, err error) map[string]interface{} {
	if err == nil {
		return data
	}
	data["error"] = err.Error()
	return data
}
