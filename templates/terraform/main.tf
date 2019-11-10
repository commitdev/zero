provider "aws" {
  region              = "${var.region}"
}

# {{ if .Config.Infrastructure.AWS.Terraform.RemoteState }}
# Store remote state in S3
resource "aws_s3_bucket" "terraform_remote_state" {
  bucket  = "${ var.remote_state_s3_bucket }"
  acl     = "private"

  versioning {
    enabled = true
  }
}

resource "aws_dynamodb_table" "terraform_state_locks" {
  name           = "${ var.remote_state_dynamo_table }"
  read_capacity  = 2
  write_capacity = 2
  hash_key       = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}

# Reference the remote state
terraform {
  backend "s3" {
    bucket         = "${var.remote_state_s3_bucket}"
    key            = "infrastructure/terraform/shared"
    encrypt        = true
    region         = "${var.region}"
    dynamodb_table = "${var.remote_state_dynamo_table}"
  }
}
# {{- end}}
# {{ if .Config.Infrastructure.AWS.Cognito }}
# ref: https://github.com/squidfunk/terraform-aws-cognito-auth#usage

# data "aws_acm_certificate" "wildcard_cert" {
#   domain   = "*.${var.public_dns_zone}"
# }

module "cognito-auth" {
  source  = "squidfunk/cognito-auth/aws"
  version = "0.4.2"

  namespace                      = "${var.auth_namespace}"
  region                         = "${var.region}"
  cognito_identity_pool_name     = "${var.auth_pool_name}"
  cognito_identity_pool_provider = "${var.auth_pool_provider}"

  # Optional: Default UI
  # app_hosted_zone_id             = "<hosted-zone-id>"
  # app_certificate_arn            = "${data.aws_acm_certificate.wildcard_cert.arn}"
  # app_domain                     = "<domain>"
  # app_origin                     = "<origin-domain>"

  # Optional: Email delivery
  # ses_sender_address             = "<email>"
}
# {{- end}}
