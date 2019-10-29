# Allow the worker nodes to assume a role we are creating below
data "aws_iam_policy_document" "k8s_worker_assumerole_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "AWS"
      identifiers = [var.eks_worker_iam_role_arn]
    }
  }
}

# Policy to allow worker nodes to assume roles starting with "k8s-"
data "aws_iam_policy_document" "node_assume_kube2iam_role" {
  statement {
    effect    = "Allow"
    actions   = ["sts:AssumeRole"]
    resources = ["arn:aws:iam::${var.iam_account_id}:role/k8s-*"]
  }
}

# Add the above policy to the worker role
resource "aws_iam_role_policy" "node_kube2iam_policy" {
  name   = "eks-node-kube2iam-policy"
  role   = var.eks_worker_iam_role_name
  policy = data.aws_iam_policy_document.node_assume_kube2iam_role.json
}

# This is now done with the kubernetes terraform provider, see the kubernetes/kube2iam module.
# # Execute the kubernetes manifest required to create the kube2iam daemonset
# resource "null_resource" "kube2iam" {
#   provisioner "local-exec" {
#     command = "kubectl apply -f ${path.root}/kubernetes/kube2iam.yaml --kubeconfig ${path.root}/output/kubeconfig_${var.environment}"
#   }
#   # TODO: Module-aware dependencies not yet supported - https://github.com/hashicorp/terraform/issues/17101
#   # depends_on = ["module.eks"]
# }

### Kube2IAM roles to map to pods ###
# These can be referenced in an annotation in a kubernetes deployment manifest file

## ALB Ingress Controller
# Create a role and establish a trust relationship with the worker nodes
resource "aws_iam_role" "k8s_worker_alb_ingress_controller_role" {
  name                  = "k8s-alb-ingress-controller"
  assume_role_policy    = data.aws_iam_policy_document.k8s_worker_assumerole_policy.json
  force_detach_policies = true
}

# Policy allowing access to specific AWS resources
data "aws_iam_policy_document" "k8s_alb_ingress_controller_access_policy" {
  statement {
    actions = [
      "acm:DescribeCertificate",
      "acm:ListCertificates",
      "acm:GetCertificate",
      "ec2:AuthorizeSecurityGroupIngress",
      "ec2:CreateSecurityGroup",
      "ec2:CreateTags",
      "ec2:DeleteTags",
      "ec2:DeleteSecurityGroup",
      "ec2:DescribeAccountAttributes",
      "ec2:DescribeInstances",
      "ec2:DescribeInstanceStatus",
      "ec2:DescribeInternetGateways",
      "ec2:DescribeSecurityGroups",
      "ec2:DescribeSubnets",
      "ec2:DescribeTags",
      "ec2:DescribeVpcs",
      "ec2:ModifyInstanceAttribute",
      "ec2:ModifyNetworkInterfaceAttribute",
      "ec2:RevokeSecurityGroupIngress",
      "elasticloadbalancing:AddTags",
      "elasticloadbalancing:CreateListener",
      "elasticloadbalancing:CreateLoadBalancer",
      "elasticloadbalancing:CreateRule",
      "elasticloadbalancing:CreateTargetGroup",
      "elasticloadbalancing:DeleteListener",
      "elasticloadbalancing:DeleteLoadBalancer",
      "elasticloadbalancing:DeleteRule",
      "elasticloadbalancing:DeleteTargetGroup",
      "elasticloadbalancing:DeregisterTargets",
      "elasticloadbalancing:DescribeListenerCertificates",
      "elasticloadbalancing:DescribeListeners",
      "elasticloadbalancing:DescribeLoadBalancers",
      "elasticloadbalancing:DescribeLoadBalancerAttributes",
      "elasticloadbalancing:DescribeRules",
      "elasticloadbalancing:DescribeSSLPolicies",
      "elasticloadbalancing:DescribeTags",
      "elasticloadbalancing:DescribeTargetGroups",
      "elasticloadbalancing:DescribeTargetGroupAttributes",
      "elasticloadbalancing:DescribeTargetHealth",
      "elasticloadbalancing:ModifyListener",
      "elasticloadbalancing:ModifyLoadBalancerAttributes",
      "elasticloadbalancing:ModifyRule",
      "elasticloadbalancing:ModifyTargetGroup",
      "elasticloadbalancing:ModifyTargetGroupAttributes",
      "elasticloadbalancing:RegisterTargets",
      "elasticloadbalancing:RemoveTags",
      "elasticloadbalancing:SetIpAddressType",
      "elasticloadbalancing:SetSecurityGroups",
      "elasticloadbalancing:SetSubnets",
      "elasticloadbalancing:SetWebACL",
      "iam:GetServerCertificate",
      "iam:ListServerCertificates",
      "waf-regional:GetWebACLForResource",
      "waf-regional:GetWebACL",
      "waf-regional:AssociateWebACL",
      "waf-regional:DisassociateWebACL",
      "waf:GetWebACL",
      "tag:GetResources",
      "tag:TagResources",
      "cognito-idp:DescribeUserPoolClient",
    ]

    resources = ["*"]
  }

  statement {
    actions   = ["iam:CreateServiceLinkedRole"]
    resources = ["arn:aws:iam::${var.iam_account_id}:role/aws-service-role/elasticloadbalancing.amazonaws.com/AWSServiceRoleForElasticLoadBalancing"]
  }
}

# Add the above policy to the created role
resource "aws_iam_role_policy" "k8s_worker_alb_ingress_controller_role_policy" {
  name   = "worker-alb-ingress-controller-policy"
  role   = aws_iam_role.k8s_worker_alb_ingress_controller_role.id
  policy = data.aws_iam_policy_document.k8s_alb_ingress_controller_access_policy.json
}