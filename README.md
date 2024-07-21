# sqs
A fully managed message queue service offered by AWS. It provides a reliable, scalable, and cost-effective way to decouple and coordinate distributed software systems and microservices.

### Libraries for Amazon SQS (Simple Queue Service)
- GO: [sqs](https://github.com/core-go/sqs), to wrap and simplify [aws-sdk-go/service/sqs](https://github.com/aws/aws-sdk-go/tree/main/service/sqs). Example is at [go-amazon-sqs-sample](https://github.com/project-samples/go-amazon-sqs-sample)

#### A common flow to consume a message from a message queue
![A common flow to consume a message from a message queue](https://cdn-images-1.medium.com/max/800/1*Y4QUN6QnfmJgaKigcNHbQA.png)
- The libraries to implement this flow are:
  - [mq](https://github.com/core-go/mq) for GOLANG. Example is at [go-amazon-sqs-sample](https://github.com/project-samples/go-amazon-sqs-sample)

### Use Cases of Amazon SQS (Simple Queue Service)
![Microservice Architecture](https://cdn-images-1.medium.com/max/800/1*vKeePO_UC73i7tfymSmYNA.png)
#### Decoupling Microservices
- <b>Scenario</b>: Separating different parts of an application to ensure that a failure in one part does not affect others.
- <b>Benefit</b>: Enhances fault tolerance and scalability by allowing asynchronous communication between services.
#### Asynchronous Processing
- <b>Scenario</b>: Handling tasks that do not need immediate processing, such as batch processing or background tasks.
- <b>Benefit</b>: Improves system efficiency and response times for end-users.
  ![A typical micro service](https://cdn-images-1.medium.com/max/800/1*d9kyekAbQYBxH-C6w38XZQ.png)
#### Job Queuing
- <b>Scenario</b>: Managing and distributing jobs to worker processes.
- <b>Benefit</b>: Balances load and ensures all tasks are completed without overloading any single worker.
#### Order Processing Systems
- <b>Scenario</b>: Processing customer orders, where each order can be handled as a separate task.
- <b>Benefit</b>: Ensures reliable and scalable processing of orders, even during high demand.
#### Message Buffering
- <b>Scenario</b>: Smoothing out bursty traffic in applications to prevent overload.
- <b>Benefit</b>: Protects the system from spikes in traffic by buffering messages.
#### Workflow Orchestration
- <b>Scenario</b>: Orchestrating steps in a complex workflow, such as image processing pipelines.
- <b>Benefit</b>: Coordinates different stages of processing in a reliable and scalable manner.

## Comparison of Amazon SQS, Google Pub/Sub and Apache Kafka
#### Amazon SQS
- <b>Type</b>: Managed message queuing service.
- <b>Use Case</b>: Decoupling and scaling microservices, asynchronous tasks.
- <b>Scalability</b>: Automatically scales.
- <b>Delivery Guarantees</b>: At-least-once, FIFO (exactly-once).
- <b>Integration</b>: Deep integration with AWS services.
- <b>Delivery Models</b>: Primarily pull, with long polling.

#### Google Pub/Sub:
- <b>Type</b>: Managed real-time messaging service.
- <b>Use Case</b>: Event-driven architectures, real-time analytics.
- <b>Scalability</b>: Automatically scales.
- <b>Delivery Guarantees</b>: At-least-once delivery.
- <b>Integration</b>: Tight with Google Cloud services.
- <b>Delivery Models</b>: Push and pull.

#### Apache Kafka
- <b>Type</b>: Open-source event streaming platform.
- <b>Use Case</b>: High-throughput messaging, event sourcing, log aggregation.
- <b>Scalability</b>: High with partitioned topics.
- <b>Delivery Guarantees</b>: Configurable (at-least-once, exactly-once).
- <b>Integration</b>: Broad ecosystem with various connectors.
- <b>Delivery Models</b>: Pull-based consumer groups.

### Key Differences
- <b>Management</b>: Pub/Sub and SQS are managed services, while Kafka is typically self-managed or via managed services like Confluent.
- <b>Use Case Focus</b>: Pub/Sub and Kafka are ideal for real-time processing, whereas SQS is great for decoupling microservices and handling asynchronous tasks.
- <b>Delivery Models</b>: Pub/Sub supports push and pull, SQS supports pull with long polling, and Kafka primarily uses pull with consumer groups.
- <b>Scalability</b>: All three are highly scalable, but Kafka offers the most control over performance tuning.
- <b>Integration</b>: Pub/Sub integrates well with Google Cloud, SQS with AWS, and Kafka has a broad integration ecosystem.

## Installation

Please make sure to initialize a Go module before installing core-go/sqs:

```shell
go get -u github.com/core-go/sqs
```

Import:

```go
import "github.com/core-go/sqs"
```
