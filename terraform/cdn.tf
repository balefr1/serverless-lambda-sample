resource "aws_cloudfront_distribution" "sample_app_website" {
  depends_on = [
    aws_s3_bucket.website_bucket
  ]

  origin {
    domain_name = aws_s3_bucket.website_bucket.bucket_regional_domain_name
    origin_id   = "s3-cloudfront"

    origin_access_control_id = aws_cloudfront_origin_access_control.website-cfront-origin-access.id
  }

  enabled             = true
  default_root_object = "index.html"


  default_cache_behavior {
    allowed_methods = [
      "GET",
      "HEAD",
    ]

    cached_methods = [
      "GET",
      "HEAD",
    ]

    target_origin_id = "s3-cloudfront"

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "redirect-to-https"
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }
  restrictions {
    geo_restriction {
        restriction_type = "none"
    }
  }
  price_class = "PriceClass_100"
}

resource "aws_cloudfront_origin_access_control" "website-cfront-origin-access" {
  name                              = "website-cfront-origin-access"
  description                       = "website-cfront-origin-access Policy"
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}