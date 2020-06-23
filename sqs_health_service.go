package sqs

import (
	"context"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSSubHealthService struct {
	Client    *sqs.SQS
	QueueName *string
	Service   string
}

func NewSQSHealthService(client *sqs.SQS, name string, queueName string) *SQSSubHealthService {
	return &SQSSubHealthService{client, &queueName, name}
}

func (h *SQSSubHealthService) Name() string {
	return h.Service
}

func (h *SQSSubHealthService) Check(ctx context.Context) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	_, err := h.Client.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: h.QueueName,
	})
	return res, err
}

func (h *SQSSubHealthService) Build(ctx context.Context, data map[string]interface{}, err error) map[string]interface{} {
	if err == nil {
		return data
	}
	data["error"] = err.Error()
	return data
}
