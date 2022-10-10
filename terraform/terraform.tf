terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.25"
    }
  }
}

provider "aws" {
  region = "eu-south-1"
}

provider "aws" {
  alias = "aws-us"
  region = "us-east-1"
}
