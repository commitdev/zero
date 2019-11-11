variable "project" {
  description = "Name of the project"
}

variable "environment" {
  description = "The environment (dev/staging/prod)"
}

variable "assume_role_policy" {
  description = "IAM policy document for AssumeRole"
}

variable "private_subnets" {
  description = "VPC subnets for the EKS cluster"
  # type        = list(string)
}

variable "vpc_id" {
  description = "VPC ID for EKS cluster"
}

variable "worker_instance_type" {
  description = "Instance type for the EKS workers"
}

variable "worker_asg_max_size" {
  description = "Maximum number of instances for the EKS ASG"
}

variable "worker_ami" {
  description = "The (EKS-optimized) AMI for EKS worker instances"
}

variable "iam_account_id" {
  description = "Account ID of the current IAM user"
}

