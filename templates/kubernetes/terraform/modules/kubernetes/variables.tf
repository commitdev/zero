variable "region" {
  description = "AWS Region"
}

variable "environment" {
  description = "Environment"
}

variable "cluster_name" {
  description = "Kubernetes cluster name"
}

variable "assume_role_policy" {
  description = "Assume-role policy for monitoring"
}