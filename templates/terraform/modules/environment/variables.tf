variable "project" {
  description = "The name of the project, mostly for tagging"
}

variable "environment" {
  description = "The environment (dev/staging/prod)"
}

variable "region" {
  description = "The AWS region"
}

variable "allowed_account_ids" {
  description = "The IDs of AWS accounts for this project, to protect against mistakenly applying to the wrong env"
  type        = list(string)
}

variable "ecr_repositories" {
  description = "List of ECR repository names to create"
  type        = list(string)
}

variable "eks_worker_instance_type" {
  description = "Instance type for the EKS workers"
}

variable "eks_worker_asg_max_size" {
  description = "Maximum number of instances for the EKS ASG"
}

variable "eks_worker_ami" {
  description = "The (EKS-optimized) AMI for EKS worker instances"
}

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
