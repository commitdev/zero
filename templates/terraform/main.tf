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
# {{ if .Config.Infrastructure.AWS.Cognito.Deploy }}
resource "aws_cognito_user_pool" "users" {
  name = "${var.project}-user-pool"

  username_attributes = [
    "email",
  ]

  auto_verified_attributes = [ "email"]
}

resource "aws_cognito_user_pool_client" "client" {
  name = "${var.user_pool}-user-pool-client"

  user_pool_id    = "${aws_cognito_user_pool.users.id}"
  generate_secret = false

  allowed_oauth_flows_user_pool_client = true
  allowed_oauth_flows = ["code", "implicit"]
  allowed_oauth_scopes = ["profile", "openid"]

  supported_identity_providers = ["COGNITO"]
  refresh_token_validity = "14"

  explicit_auth_flows = [
    "ADMIN_NO_SRP_AUTH",
    "USER_PASSWORD_AUTH",
  ]

  write_attributes = ["email"]

  # TODO : Vars for this subdomain
  callback_urls = ["https://auth.${var.hostname}","https://auth.${var.hostname}/oauth2/idpresponse"]
  logout_urls = ["https://auth.${var.hostname}/logout"]
}

output "cognito_pool_id" {
  value = "${aws_cognito_user_pool.users.id}"
}
output "cognito_client_id" {
  value = "${aws_cognito_user_pool_client.client.id}"
}
# {{- end}}
