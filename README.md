# FarmQ

FarmQ is a distributed task queue where clients submit jobs to be executed asynchronously by worker nodes.

## Features

- Asynchronous job submission
- Distributed worker processing
- Retry handling
- Dead-letter queue support
- Result storage in S3
- Job metadata stored in DynamoDB
- Containerized services
- CI/CD with GitHub Actions

## High Level Architecture

Client → API Service → SQS service → Worker Service → S3 + DynamoDB

## Tech Stack

- Go
- AWS SQS
- AWS ECS (Fargate)
- DynamoDB
- S3
- Docker
- GitHub Actions
- Terraform
