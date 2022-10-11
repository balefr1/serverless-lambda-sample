// API GW

resource "aws_api_gateway_rest_api" "sample-api" {
  name = "sample-api"
  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_api_gateway_rest_api_policy" "sample-api" {
  count = 0
  rest_api_id = aws_api_gateway_rest_api.sample-api.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "*"
      },
      "Action": "execute-api:Invoke",
      "Resource": "${format("%s/%s/%s/%s",aws_api_gateway_rest_api.sample-api.execution_arn,"*","*","*")}",
      "Condition": {
        "IpAddress": {
          "aws:SourceIp": ["130.25.92.10/32"]
        }
      }
    }
  ]
}
EOF
}


resource "aws_api_gateway_resource" "sample-api-proxy" {
  rest_api_id = aws_api_gateway_rest_api.sample-api.id
  parent_id   = aws_api_gateway_rest_api.sample-api.root_resource_id
  path_part   = "{proxy+}"
}

resource "aws_api_gateway_method" "sample-api-proxy" {
  rest_api_id   = aws_api_gateway_rest_api.sample-api.id
  resource_id   = aws_api_gateway_resource.sample-api-proxy.id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "sample-api-proxy_integration" {
  rest_api_id             = aws_api_gateway_rest_api.sample-api.id
  resource_id             = aws_api_gateway_resource.sample-api-proxy.id
  http_method             = aws_api_gateway_method.sample-api-proxy.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_alias.sample-api_alias.invoke_arn
}

resource "aws_api_gateway_deployment" "sample-api-deploy" {
  rest_api_id = aws_api_gateway_rest_api.sample-api.id
  triggers = {
    redeployment = sha1(jsonencode([
      aws_api_gateway_resource.sample-api-proxy.id,
      aws_api_gateway_method.sample-api-proxy.id,
      aws_api_gateway_integration.sample-api-proxy_integration.id
      #aws_api_gateway_rest_api_policy.sample-api.id
    ]))
  }
}
resource "aws_api_gateway_stage" "sample-api-stage" {
  deployment_id = aws_api_gateway_deployment.sample-api-deploy.id
  rest_api_id   = aws_api_gateway_rest_api.sample-api.id
  stage_name    = "dev"
  lifecycle {
    ignore_changes = [deployment_id]
  }
}
