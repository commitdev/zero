variable "region" {
  description = "The AWS region"
}

variable "remote_state_s3_bucket" {
  description = "Name of the S3 bucket to store the remote state"
}

variable "remote_state_dynamo_table" {
  description = "Dynamo DB Table to store the remote state locks"
}