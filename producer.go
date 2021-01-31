package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Producer struct {
	Client       *sqs.SQS
	QueueURL     *string
	DelaySeconds *int64 //could be 10
}

func NewProducerByQueueName(client *sqs.SQS, queueName string, delaySeconds *int64) (*Producer, error) {
	queueUrl, err := GetQueueUrl(client, queueName)
	if err != nil {
		return nil, err
	}
	return NewProducer(client, queueUrl, delaySeconds), nil
}

func NewProducer(client *sqs.SQS, queueURL string, delaySeconds *int64) *Producer {
	return &Producer{Client: client, QueueURL: &queueURL, DelaySeconds: delaySeconds}
}

func (p *Producer) Produce(ctx context.Context, data []byte, messageAttributes map[string]string) (string, error) {
	attributes := MapToAttributes(messageAttributes)
	s := string(data)
	result, err := p.Client.SendMessage(&sqs.SendMessageInput{
		DelaySeconds:      p.DelaySeconds,
		MessageAttributes: attributes,
		MessageBody:       aws.String(s),
		QueueUrl:          p.QueueURL,
	})
	if result != nil && result.MessageId != nil {
		return *result.MessageId, err
	} else {
		return "", err
	}
}
