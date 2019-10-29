variable "environment" {
  description = "The environment (dev/staging/prod)"
}

variable "ecr_repositories" {
  description = "List of ECR repository names to create"
  type        = list(string)
}

