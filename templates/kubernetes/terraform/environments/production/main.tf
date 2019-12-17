terraform {
  backend "s3" {
    bucket         = "{{ .Config.Name }}-production-terraform-state"
    key            = "infrastructure/terraform/environments/production/kubernetes"
    encrypt        = true
    region         = "{{ .Config.Infrastructure.AWS.Region }}"
    dynamodb_table = "{{ .Config.Name }}-production-terraform-state-locks"
  }
}

# Provision kubernetes resources required to run services/applications
module "kubernetes" {
  source = "../../modules/kubernetes"

  environment = "production"
  region      = "{{ .Config.Infrastructure.AWS.Region }}"

  # Authenticate with the EKS cluster via the cluster id
  cluster_name = "{{ .Config.Infrastructure.AWS.EKS.ClusterName }}"

  # Assume-role policy used by monitoring fluentd daemonset
  assume_role_policy = data.aws_iam_policy_document.assumerole_root_policy.json

  external_dns_zone = "{{ .Config.Frontend.Hostname }}"
  external_dns_owner_id = "{{ GenerateUUID }}"
  external_dns_assume_roles = [ "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/k8s-{{ .Config.Infrastructure.AWS.EKS.ClusterName }}-workers" ]
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
