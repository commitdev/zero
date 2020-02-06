locals {
  assets_access_identity = "${var.project}-client-assets"
}

resource "aws_s3_bucket" "client_assets" {
  for_each = var.buckets

  // Our bucket's name is going to be the same as our site's domain name.
  bucket = each.value
  acl    = "private" // The contents will be available through cloudfront, they should not be accessible publicly


  // S3 understands what it means to host a website.
  website {
    // Here we tell S3 what to use when a request comes in to the root
    index_document = "index.html"
    error_document = "404.html"
  }
}

# Deny public access to this bucket
resource "aws_s3_bucket_public_access_block" "client_assets" {
  for_each = var.buckets

  bucket                  = each.value
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# Access identity for CF access to S3
resource "aws_cloudfront_origin_access_identity" "client_assets" {
  comment = local.assets_access_identity
}

# Policy to allow CF access to S3
data "aws_iam_policy_document" "assets_origin" {
  for_each = var.buckets

  statement {
    actions   = ["s3:GetObject"]
    resources = ["arn:aws:s3:::${each.value}/*"]

    principals {
      type        = "AWS"
      identifiers = [aws_cloudfront_origin_access_identity.client_assets.iam_arn]
    }
  }

  statement {
    actions   = ["s3:ListBucket"]
    resources = ["arn:aws:s3:::${each.value}"]

    principals {
      type        = "AWS"
      identifiers = [aws_cloudfront_origin_access_identity.client_assets.iam_arn]
    }
  }
}

# Attach the policy to the bucket
resource "aws_s3_bucket_policy" "client_assets" {
  for_each = var.buckets

  bucket = each.value
  policy = data.aws_iam_policy_document.assets_origin[each.value].json
}

# To use an ACM cert with CF it has to exist in us-east-1
provider "aws" {
  region = "us-east-1"
  alias  = "east1"
}

# Find an already created ACM cert for this domain
data "aws_acm_certificate" "wildcard_cert" {
  provider    = "aws.east1"
  domain      = var.cert_domain
  most_recent = "true"
}

# Create the cloudfront distribution
resource "aws_cloudfront_distribution" "client_assets_distribution" {
  for_each = var.buckets

  // origin is where CloudFront gets its content from.
  origin {
      domain_name = aws_s3_bucket.client_assets[each.value].bucket_domain_name
      origin_id   = local.assets_access_identity
      s3_origin_config {
        origin_access_identity = aws_cloudfront_origin_access_identity.client_assets.cloudfront_access_identity_path
      }
    }

  enabled             = true
  is_ipv6_enabled     = true
  default_root_object = "index.html" # Render this when you hit the root

  // All values are defaults from the AWS console.
  default_cache_behavior {
    target_origin_id       = local.assets_access_identity
    viewer_protocol_policy = "redirect-to-https"
    compress               = true
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]
    min_ttl                = 0
    default_ttl            = 86400
    max_ttl                = 31536000

    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }
  }

  aliases = [
    each.value,
  ]

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  # Use our cert
  viewer_certificate {
      acm_certificate_arn      = data.aws_acm_certificate.wildcard_cert.arn
      minimum_protocol_version = "TLSv1"
      ssl_support_method       = "sni-only"
    }

}

# Find the route53 zone
data "aws_route53_zone" "public" {
  name         = "${var.cert_domain}."
  private_zone = false
}

# Subdomain to point at CF
resource "aws_route53_record" "client_assets" {
  for_each = var.buckets

  zone_id = data.aws_route53_zone.public.zone_id
  name    = each.value
  type    = "CNAME"
  ttl     = "120"
  records = [aws_cloudfront_distribution.client_assets_distribution[each.value].domain_name]
}
