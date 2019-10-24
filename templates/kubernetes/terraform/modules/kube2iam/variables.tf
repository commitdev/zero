variable "eks_worker_iam_role_arn" {
  description = "The ARN of the EKS worker IAM role"
}

variable "eks_worker_iam_role_name" {
  description = "The name of the EKS worker IAM role"
}

variable "iam_account_id" {
  description = "Account ID of the current IAM user"
}

variable "environment" {
  description = "The environment (dev/staging/prod)"
}