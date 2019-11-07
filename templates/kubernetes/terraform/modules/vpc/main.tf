module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = "${var.project}-${var.environment}-vpc"
  cidr = "10.20.0.0/16"

  azs              = ["${var.region}a", "${var.region}b", "${var.region}c"] # Most regions have 3+ azs
  private_subnets  = ["10.20.40.0/24", "10.20.42.0/24", "10.20.44.0/24"]
  public_subnets   = ["10.20.41.0/24", "10.20.43.0/24", "10.20.45.0/24"]
  database_subnets = ["10.20.50.0/24", "10.20.52.0/24", "10.20.54.0/24"]

  # Allow kubernetes ALB ingress controller to auto-detect
  private_subnet_tags = {
    "kubernetes.io/cluster/${var.project}-${var.environment}" = "owned"
    "kubernetes.io/role/internal-elb"                         = "1"
  }

  public_subnet_tags = {
    "kubernetes.io/cluster/${var.project}-${var.environment}" = "owned"
    "kubernetes.io/role/elb"                                  = "1"
  }

  enable_nat_gateway   = true
  enable_vpn_gateway   = false
  enable_dns_hostnames = true

  create_database_subnet_group       = true
  create_database_subnet_route_table = true

  tags = {
    environment = var.environment
  }

}
