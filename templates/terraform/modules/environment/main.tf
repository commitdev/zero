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

# @TODO - Move this to a different file

# {{ if .Config.Infrastructure.AWS.Cognito }}
# ref: https://github.com/squidfunk/terraform-aws-cognito-auth#usage

# data "aws_acm_certificate" "wildcard_cert" {
#   domain   = "*.${var.public_dns_zone}"
# }

module "cognito-auth" {
  source  = "squidfunk/cognito-auth/aws"
  version = "0.4.2"

  namespace                      = "${var.auth_namespace}"
  region                         = "${var.region}"
  cognito_identity_pool_name     = "${var.auth_pool_name}"
  cognito_identity_pool_provider = "${var.auth_pool_provider}"

  # Optional: Default UI
  # app_hosted_zone_id             = "<hosted-zone-id>"
  # app_certificate_arn            = "${data.aws_acm_certificate.wildcard_cert.arn}"
  # app_domain                     = "<domain>"
  # app_origin                     = "<origin-domain>"

  # Optional: Email delivery
  # ses_sender_address             = "<email>"
}
# {{- end}}
