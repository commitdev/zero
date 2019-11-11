# Environment entrypoint

module "vpc" {
  source      = "../../modules/vpc"
  project     = var.project
  environment = var.environment
  region      = var.region
}

# Data sources for EKS IAM
data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "assumerole_root_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
    }
  }
}

# Provision the EKS cluster
module "eks" {
  source               = "../../modules/eks"
  project              = var.project
  environment          = var.environment
  assume_role_policy   = data.aws_iam_policy_document.assumerole_root_policy.json
  private_subnets      = module.vpc.private_subnets
  vpc_id               = module.vpc.vpc_id
  worker_instance_type = var.eks_worker_instance_type
  worker_asg_max_size  = var.eks_worker_asg_max_size
  worker_ami           = var.eks_worker_ami # EKS-Optimized AMI for your region: https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html
  iam_account_id       = data.aws_caller_identity.current.account_id
}

module "kube2iam" {
  source                   = "../../modules/kube2iam"
  environment              = var.environment
  eks_worker_iam_role_arn  = module.eks.worker_iam_role_arn
  eks_worker_iam_role_name = module.eks.worker_iam_role_name
  iam_account_id           = data.aws_caller_identity.current.account_id
}

# {{ if .Config.Infrastructure.AWS.Cognito.Enabled }}
resource "cognito" "auth" {
  user_pool   = var.user_pool
  hostname    = var.hostname
}
# {{- end}}

# {{ if .Config.Infrastructure.AWS.S3Hosting.Enabled }}
resource "s3_hosting" "assets" {
  bucket_name   = var.s3_hosting_bucket_name
}
# {{- end}}
