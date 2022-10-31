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

Notes:
- after creating cloudfront distribution and s3 bucket for website, frontend (react) must be manually built and uploaded to s3.
    In order to build react, go to frontend/react dir and run "npm run build"
    Upload to s3 can be done using sed_sync script, specifying bucket name and cloudfront distribution fqdn.