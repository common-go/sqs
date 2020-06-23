package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/common-go/mq"
)

type Consumer struct {
	Client            *sqs.SQS
	QueueURL          *string
	AckOnConsume      bool
	VisibilityTimeout int64 // should be 20 (seconds)
	WaitTimeSeconds   int64 // should be 0
}

func NewConsumerByQueueName(client *sqs.SQS, queueName string, ackOnConsume bool, visibilityTimeout int64, waitTimeSeconds int64) (*Consumer, error) {
	queueUrl, err := GetQueueUrl(client, queueName)
	if err != nil {
		return nil, err
	}
	return NewConsumer(client, queueUrl, ackOnConsume, visibilityTimeout, waitTimeSeconds), nil
}

func NewConsumer(client *sqs.SQS, queueURL string, ackOnConsume bool, visibilityTimeout int64, waitTimeSeconds int64) *Consumer {
	return &Consumer{Client: client, QueueURL: &queueURL, AckOnConsume: ackOnConsume, VisibilityTimeout: visibilityTimeout, WaitTimeSeconds: waitTimeSeconds}
}

func (c *Consumer) Consume(ctx context.Context, caller mq.ConsumerCaller) {
	result, er1 := c.Client.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            c.QueueURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(c.VisibilityTimeout), // 20 seconds
		WaitTimeSeconds:     aws.Int64(c.WaitTimeSeconds),
	})
	if er1 != nil {
		caller.Call(ctx, nil, er1)
	} else {
		if len(result.Messages) > 0 {
			m := result.Messages[0]
			data := []byte(*m.Body)
			attributes := PtrToMap(m.Attributes)
			message := mq.Message{
				Id:         *m.MessageId,
				Data:       data,
				Attributes: attributes,
				Raw:        m,
			}
			if c.AckOnConsume {
				_, er2 := c.Client.DeleteMessage(&sqs.DeleteMessageInput{
					QueueUrl:      c.QueueURL,
					ReceiptHandle: result.Messages[0].ReceiptHandle,
				})
				if er2 != nil {
					caller.Call(ctx, nil, er2)
				} else {
					caller.Call(ctx, &message, nil)
				}
			} else {
				caller.Call(ctx, &message, nil)
			}
		}
	}
}

func PtrToMap(m map[string]*string) map[string]string {
	attributes := make(map[string]string)
	for k, v := range m {
		if v != nil {
			attributes[k] = *v
		}
	}
	return attributes
}
