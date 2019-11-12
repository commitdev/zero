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

# {{ if .Config.Infrastructure.AWS.Cognito.Deploy }}
resource "aws_cognito_user_pool" "users" {
  name = "${var.user_pool}-user-pool"

  username_attributes = [
    "email",
  ]

  # auto_verified_attributes = ["email"]
}

resource "aws_cognito_user_pool_client" "client" {
  name = "${var.user_pool}-cognito-client"

  user_pool_id    = "${aws_cognito_user_pool.users.id}"
  generate_secret = false

  allowed_oauth_flows_user_pool_client = true
  allowed_oauth_flows = ["code", "implicit"]
  allowed_oauth_scopes = ["profile", "openid"]

  supported_identity_providers = ["COGNITO"]
  refresh_token_validity = "14"

  explicit_auth_flows = [
    "ADMIN_NO_SRP_AUTH",
    "USER_PASSWORD_AUTH",
  ]

  write_attributes = ["email"]

  callback_urls = ["https://auth.${var.hostname}","https://auth.${var.hostname}/oauth2/idpresponse"]
  logout_urls = ["https://auth.${var.hostname}/logout"]
}

output "cognito_pool_id" {
  value = "${aws_cognito_user_pool.users.id}"
}
output "cognito_client_id" {
  value = "${aws_cognito_user_pool_client.client.id}"
}
# {{- end}}
