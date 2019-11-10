variable "region" {
  default = "{{ .Config.Infrastructure.AWS.Region }}"
  description = "The AWS region"
}

# {{ if .Config.Infrastructure.AWS.Terraform.RemoteState }}
variable "remote_state_s3_bucket" {
  default = "project-{{ .Config.Name }}-terraform-state"
  description = "Name of the S3 bucket to store the remote state"
}

variable "remote_state_dynamo_table" {
  default = "{{ .Config.Name }}-terraform-state-locks"
  description = "Dynamo DB Table to store the remote state locks"
}
# {{- end}}
# {{ if .Config.Infrastructure.AWS.Cognito }}
variable "auth_namespace" {
  default = "cognito_auth"
} 
variable "auth_pool_name" {
  description = "AWS Cognito pool name"
} 
variable "auth_pool_provider" {
  description = "AWS Cognito pool provider"
}
# {{- end}}
