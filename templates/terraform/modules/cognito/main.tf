resource "aws_cognito_user_pool" "users" {
  name = "${var.user_pool}-user-pool"

  username_attributes = [
    "email",
  ]

  # auto_verified_attributes = ["email"]
}

resource "aws_cognito_user_pool_client" "client" {
  name = "${var.user_pool}-cognito-client"

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

  callback_urls = ["https://auth.${var.hostname}","https://auth.${var.hostname}/oauth2/idpresponse"]
  logout_urls = ["https://auth.${var.hostname}/logout"]
}

output "cognito_pool_id" {
  value = aws_cognito_user_pool.users.id
}
output "cognito_client_id" {
  value = aws_cognito_user_pool_client.client.id
}
