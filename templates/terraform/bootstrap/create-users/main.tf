provider "aws" {
  region  = "{{ .Config.Infrastructure.AWS.Region }}"
}

# Create the CI User
resource "aws_iam_user" "ci_user" {
  name = "ci-user"
}

# Create a keypair to be used by CI systems
resource "aws_iam_access_key" "ci_user" {
  user    = aws_iam_user.ci_user.name
}

# Add the keys to AWS secrets manager
resource "aws_secretsmanager_secret" "ci_user_keys" {
  name = "ci-user-keys"
}

resource "aws_secretsmanager_secret_version" "ci_user_keys" {
  secret_id     = aws_secretsmanager_secret.ci_user_keys.id
  secret_string = jsonencode(map("access_key_id", aws_iam_access_key.ci_user.id, "secret_key", aws_iam_access_key.ci_user.secret))
}
