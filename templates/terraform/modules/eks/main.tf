# Create KubernetesAdmin role for aws-iam-authenticator
resource "aws_iam_role" "kubernetes_admin_role" {
  name               = "{{ .Config.Name }}-kubernetes-admin"
  assume_role_policy = var.assume_role_policy
  description        = "Kubernetes administrator role (for AWS IAM Authenticator)"
}

module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "6.0.2"

  cluster_name    = var.cluster_name
  cluster_version = "1.14"
  subnets         = var.private_subnets
  vpc_id          = var.vpc_id

  worker_groups = [
    {
      instance_type         = var.worker_instance_type
      asg_min_size          = var.worker_asg_min_size
      asg_desired_capacity  = var.worker_asg_min_size
      asg_max_size          = var.worker_asg_max_size
      ami_id                = var.worker_ami
      tags = [{
        key                 = "environment"
        value               = var.environment
        propagate_at_launch = true
      }]
    },
  ]

  map_roles = [
    {
      rolearn  = "arn:aws:iam::${var.iam_account_id}:role/kubernetes-admin"
      username = "kubernetes-admin"
      groups   = ["system:masters"]
    },
  ]
  cluster_iam_role_name = "k8s-${var.cluster_name}-cluster"
  workers_role_name = "k8s-${var.cluster_name}-workers"

  # TODO, determine if this should be true/false
  manage_aws_auth = true

  write_kubeconfig      = false
  write_aws_auth_config = false

  tags = {
    environment = var.environment
  }
}
