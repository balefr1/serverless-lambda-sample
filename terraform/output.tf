output "apigw_endpoint" {
    value = aws_api_gateway_stage.sample-api-stage.invoke_url
}

output "cfront_endpoint" {
    value = "https://${aws_cloudfront_distribution.sample_app_website.domain_name}"
}