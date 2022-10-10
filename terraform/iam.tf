resource "aws_iam_role" "sample-api-lambda" {
  name = "sample-api-lambda"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Sid    = ""
      Principal = {
        Service = "lambda.amazonaws.com"
      }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "sample-api-lambda-policy" {
  role       = aws_iam_role.sample-api-lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_policy" "sample-api-dynamo-policy" {
  name        = "sample-api-dynamo-policy"
  path        = "/"
  description = "dynamo policy for sample app api"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "dynamodb:Get*",
          "dynamodb:Describe*",
          "dynamodb:PutItem",
          "dynamodb:UpdateItem",
          "dynamodb:Scan",
          "dynamodb:Query",
          "dynamodb:UpdateTimeToLive",
          "dynamodb:DeleteItem"
        ]
        Effect   = "Allow"
        Resource = aws_dynamodb_table.sampleapi-table.arn
      },
    ]
  })
}

resource "aws_iam_policy" "sample-api-s3-policy" {
  name        = "sample-api-s3-policy"
  path        = "/"
  description = "s3 policy for sample app api"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "s3:PutObject",
          "s3:GetObject",
          "s3:HeadObject"
        ]
        Effect   = "Allow"
        Resource = [format("arn:aws:s3:::%s/*",aws_s3_bucket.sample_app_bucket.id)]
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "sample-api-lambda-db-policy" {
  role       = aws_iam_role.sample-api-lambda.name
  policy_arn = aws_iam_policy.sample-api-dynamo-policy.arn
}

resource "aws_iam_role_policy_attachment" "sample-api-lambda-s3-policy" {
  role       = aws_iam_role.sample-api-lambda.name
  policy_arn = aws_iam_policy.sample-api-s3-policy.arn
}



