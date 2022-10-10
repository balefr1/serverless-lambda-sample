Sample serverless website build on AWS cloud.

![Alt text](docs/sampleapp.jpg?raw=true "Architecture")

This example contains a React js frontend and a Golang backend running on AWS Lambda.
The simple website allows users to upload/retrieve files, which are stored in S3 and their metadata on DynamoDB.

Requires:
- Go 1.18
- Npm
- Docker
- AWS SAM cli
- Terraform > 1.2

Terraform is used to deploy all cloud resources as per shown architecture.