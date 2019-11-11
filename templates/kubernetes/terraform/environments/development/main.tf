terraform {
  backend "s3" {
    bucket         = "project-{{ .Config.Name }}-terraform-state"
    key            = "infrastructure/terraform/environments/development/kubernetes"
    encrypt        = true
    region         = "{{ .Config.Infrastructure.AWS.Region }}"
    dynamodb_table = "{{ .Config.Name }}-terraform-state-locks"
  }
}

# Provision kubernetes resources required to run services/applications
module "kubernetes" {
  source = "../../modules/kubernetes"

  environment = "development"
  region      = "{{ .Config.Infrastructure.AWS.Region }}"

  # Authenticate with the EKS cluster via the cluster id
  cluster_name = "{{ .Config.Infrastructure.AWS.EKS.ClusterName }}"

  # Assume-role policy used by monitoring fluentd daemonset
  assume_role_policy = data.aws_iam_policy_document.assumerole_root_policy.json
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
