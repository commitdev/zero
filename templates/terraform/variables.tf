variable "project" {
  default = "{{ .Config.Name }}"
  description = "The name of the project, mostly for tagging"
}

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
# {{ if .Config.Infrastructure.AWS.Cognito.Deploy }}
variable "user_pool" {
  default = "{{ .Config.Name }}"
  description = "AWS Cognito pool name"
} 
variable "hostname" {
  default = "{{ .Config.Frontend.Hostname }}"
  description = "AWS Cognito pool name"
} 
# {{- end}}
