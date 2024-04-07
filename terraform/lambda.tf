##lambda
resource "aws_lambda_function" "sample-api" {
  function_name    = "sample-api"
  runtime          = "provided.al2"
  handler          = "main"
  role             = aws_iam_role.sample-api-lambda.arn

  timeout = 15

  filename         = "./files/sample_api_lambda_wsize.zip"
  source_code_hash = filebase64sha256("./files/sample_api_lambda_wsize.zip")
  environment {
    variables = {
      S3_BUCKET = aws_s3_bucket.sample_app_bucket.id
      S3_DIR = "/user_uploads"
      DYNAMO_TABLE = aws_dynamodb_table.sampleapi-table.name
    }
  }
}

resource "aws_lambda_alias" "sample-api_alias" {
  name             = "LIVE"
  description      = "live alias"
  function_name    = aws_lambda_function.sample-api.arn
  function_version = "$LATEST"
}

resource "aws_cloudwatch_log_group" "sample-api-lambda-log" {
  name              = "/aws/lambda/sample-api"
  retention_in_days = 30
}

resource "aws_lambda_permission" "sample-api-lambda-invoke" {
  statement_id  = "sample-api-AllowExecAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.sample-api.function_name
  qualifier = "LIVE"
  principal     = "apigateway.amazonaws.com"
  source_arn    = format("%s/%s/%s/%s",aws_api_gateway_rest_api.sample-api.execution_arn,"*","*","*")
} 