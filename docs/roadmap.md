# FarmQ Roadmap

This document outlines the planned features and scope for FarmQ.  
The project is divided into **V1 (Minimum Viable Version)** and later improvements (V2+).

---

## V1 – Minimum Viable Product

The goal of V1 is to build a fully working end-to-end system that allows clients to submit jobs, have them processed asynchronously by worker nodes, and retrieve results.

### Features

- **Job Submission**
  - Clients can submit a job via `POST /jobs`
  - Only a single job type is supported initially (`uppercase` text transformation)
- **Job Processing**
  - A worker picks up jobs from the queue
  - Executes the job logic
- **Job Status Tracking**
  - Clients can check the status of a job via `GET /jobs/{id}`
  - Status values:
    - `PENDING` → waiting to be processed
    - `RUNNING` → being processed
    - `COMPLETED` → successfully processed
    - `FAILED` → failed (optional for initial testing)
- **Result Storage**
  - Results are stored in database
- **Queue**
  - Jobs are placed in a queue
- **Data Storage**
  - Job metadata stored in database
- **Dockerized Services**
  - API service and Worker service containerized with Docker
- **CI/CD**
  - GitHub Actions pipeline builds and tests the services

---

### V1 Implementation Tasks

1. **Define Job model and Status enum**
   - Create `Job` struct with fields: ID, Type, Status, Payload, CreatedAt, UpdatedAt
   - Define Status enum: PENDING, RUNNING, COMPLETED, FAILED

2. **Implement POST /jobs endpoint**
   - Accept job submission
   - Validate payload
   - Create Job record
   - Add job to queue
   - Return job ID and initial status

3. **Implement GET /jobs/{id} endpoint**
   - Return job status
   - Include result

4. **Implement in-memory queue (local testing)**
   - Simple Go channel or slice to hold pending jobs

5. **Implement Worker service (local version)**
   - Poll queue
   - Pick job
   - Store result
   - Update job status

6. **Integrate basic storage**
   - Start with in-memory storage for V1

7. **Containerize services**
   - Create Dockerfile for API service
   - Create Dockerfile for Worker service

8. **Set up GitHub Actions pipeline**
   - Run linting
   - Run unit tests
   - Build Docker images

9. **Test end-to-end**
   - Submit job via API
   - Verify worker picks it up
   - Check status updates
   - Retrieve result

---

## Future Improvements (V2+)

- Add multiple job types
- Add worker autoscaling based on queue depth
- Implement retry logic with exponential backoff
- Add dead-letter queue handling
- Add observability: structured logging, metrics, alerts
- Add Terraform deployment for infrastructure
- Add web dashboard for job management
- Add authentication and rate limiting
