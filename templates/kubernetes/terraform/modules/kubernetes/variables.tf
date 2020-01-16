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

variable "external_dns_zone" {
  description = "R53 zone that external-dns will have access to"
}

variable "external_dns_owner_id" {
  description = "Unique id of the TXT record that external-dns will use to store state (can just be a uuid)"
}

variable "external_dns_assume_roles" {
  description = "List of roles that should be able to assume the external dns role (most likely the role of the cluster worker nodes)"
  type        = list(string)
}
