variable "environment" {
  description = "The environment (dev/staging/prod)"
}

variable "ecr_repositories" {
  description = "List of ECR repository names to create"
  type        = list(string)
}

variable "ecr_principals" {
  description = "List of principals (most likely users) to give full access to the created ECR repositories"
  type        = list(string)
}

