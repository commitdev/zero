output "vpc_id" {
  description = "The ID of the created VPC"
  value       = module.vpc.vpc_id
}

output "vpc_cidr_block" {
  description = "The CIDR block of the VPC"
  value       = module.vpc.vpc_cidr_block
}

output "azs" {
  description = "Availability zones for the VPC"
  value       = module.vpc.azs
}

output "private_subnets" {
  description = "List of private subnets"
  value       = module.vpc.private_subnets
}

output "public_subnets" {
  description = "List of public subnets"
  value       = module.vpc.public_subnets
}

output "database_subnets" {
  description = "List of public subnets"
  value       = module.vpc.database_subnets
}

output "database_subnet_group" {
  description = "List of subnet groups"
  value       = module.vpc.database_subnet_group
}

