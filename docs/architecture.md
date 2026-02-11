# Architecture

## Job Lifecycle

1. Client sends POST /jobs
2. API validates input
3. Job record written to DynamoDB
4. Message published to SQS
5. Worker polls queue
6. Worker marks job RUNNING
7. Worker executes task
8. Worker writes result to S3
9. Worker updates job status

## Diagram

Client → API Service → SQS service → Worker Service → S3 + DynamoDB

## Component Responsibilities

### API Service

- Accept job submissions
- Validate input
- Persist job metadata
- Publish job to SQS
- Expose job status endpoint

### Worker Service

- Poll SQS for messages
- Claim and process jobs
- Update job status
- Store results in S3
- Handle retries and failures

### SQS

- Durable job queue
- At-least-once message delivery
- Dead-letter queue for failed jobs

### DynamoDB

- Persistent job metadata
- Job status tracking
- Retry count tracking

### S3

- Store job results and artifacts

## Delivery Guarantees

FarmQ uses SQS Standard queues, which provide at-least-once delivery.
Workers must be idempotent because duplicate messages may occur.

## Consistency Considerations

If publishing to SQS fails after writing to DynamoDB,
the API should either:

1. Retry publishing
2. Mark job as FAILED
3. Or use a transactional outbox pattern (future improvement)
