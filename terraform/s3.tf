resource "aws_s3_bucket" "sample_app_bucket" {
    bucket = "sorint-serverless-demo-userfiles"
}

resource "aws_s3_bucket" "website_bucket" {
	provider = aws.aws-us
  bucket = "sorint-serverless-demo-website"
}

resource "aws_s3_bucket_acl" "website_bucket" {
	provider = aws.aws-us
  bucket = aws_s3_bucket.website_bucket.id
  acl    = "private"
}

resource "aws_s3_bucket_policy" "allow_access_from_website_cfont" {
	provider = aws.aws-us
  bucket = aws_s3_bucket.website_bucket.id
  policy = data.aws_iam_policy_document.allow_access_from_website_cfont.json
}

data "aws_iam_policy_document" "allow_access_from_website_cfont" {
  statement {
    principals {
      type        = "Service"
      identifiers = ["cloudfront.amazonaws.com"]
    }

    actions = [
      "s3:GetObject",
    ]

    resources = [
      "${aws_s3_bucket.website_bucket.arn}/*",
    ]

    condition {
        test     = "StringEquals"
        variable = "AWS:SourceArn"
        values = [
            aws_cloudfront_distribution.sample_app_website.arn
        ]
    }
  }
}

resource "aws_s3_bucket_object" "index" {
	provider = aws.aws-us
  bucket = aws_s3_bucket.website_bucket.id
  key    = "index.html"
  source = "../sample_app/frontend/index.html"
	content_type = "text/html"
}
