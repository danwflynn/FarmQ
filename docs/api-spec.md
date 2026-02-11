# API Specification

## POST /jobs

Request:

```json
{
  "job_type": "example",
  "payload": {}
}
```

Response:

```json
{
  "job_id": "uuid",
  "status": "PENDING"
}
```

## GET /jobs/{id}

```json
{
  "job_id": "...",
  "status": "RUNNING",
  "result_location": "s3://..."
}
```

## Job Status Values

- PENDING → Job created but not yet processed
- RUNNING → Worker is actively processing
- COMPLETED → Job finished successfully
- FAILED → Job failed after retries
- RETRYING → Job failed but will be retried

## Error Responses

All errors follow this format:

```json
{
  "error": "string",
  "message": "human readable description"
}
```

### Examples

400 Bad Request

```json
{
  "error": "invalid_payload",
  "message": "job_type is required"
}
```

404 Not Found

```json
{
  "error": "job_not_found",
  "message": "No job exists with id ..."
}
```

500 Internal Server Error

```json
{
  "error": "internal_error",
  "message": "Unexpected server error"
}
```

## Validation Rules

- job_type must be non-empty string
- payload must be valid JSON
- Maximum payload size: (to be defined)
- job_type must match a supported job handler
