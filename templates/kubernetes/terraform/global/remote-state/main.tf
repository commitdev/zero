provider "aws" {
  region  = "{{ .Config.Infrastructure.AWS.Region }}"
}

resource "aws_s3_bucket" "terraform_remote_state" {
  bucket  = "project-{{ .Config.Name }}-terraform-state"
  acl     = "private"

  versioning {
    enabled = true
  }
}

resource "aws_dynamodb_table" "terraform_state_locks" {
  name           = "{{ .Config.Name }}-terraform-state-locks"
  read_capacity  = 2
  write_capacity = 2
  hash_key       = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}
