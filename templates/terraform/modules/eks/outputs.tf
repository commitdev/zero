output "cluster_id" {
  description = "Identifier of the EKS cluster"
  value       = module.eks.cluster_id
}

output "worker_iam_role_arn" {
  description = "The ARN of the EKS worker IAM role"
  value       = module.eks.worker_iam_role_arn
}

output "worker_iam_role_name" {
  description = "The name of the EKS worker IAM role"
  value       = module.eks.worker_iam_role_name
}

output "worker_security_group_id" {
  description = "The security group of the EKS worker"
  value       = module.eks.worker_security_group_id
}