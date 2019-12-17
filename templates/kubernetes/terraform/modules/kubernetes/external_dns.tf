# Trust relationship
data "aws_iam_policy_document" "external_dns_trust_relationship" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }

  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "AWS"
      identifiers = var.external_dns_assume_roles
    }
  }
}

# external-dns role
resource "aws_iam_role" "external_dns_role" {
  name = "k8s-external-dns-role"
  assume_role_policy = data.aws_iam_policy_document.external_dns_trust_relationship.json
}

data "aws_iam_policy_document" "external_dns_policy_doc" {
  statement {
    sid    = "k8sExternalDnsRead"
    effect = "Allow"

    actions = [
      "route53:ListHostedZones",
      "route53:ListResourceRecordSets",
    ]

    resources = ["*"]
  }

  statement {
    sid    = "k8sExternalDnsWrite"
    effect = "Allow"

    actions = ["route53:ChangeResourceRecordSets"]

    resources = ["arn:aws:route53:::hostedzone/*"]
  }
}

resource "aws_iam_role_policy" "external_dns_policy" {
  name = "k8s-external-dns-policy"
  role = aws_iam_role.external_dns_role.id
  policy = data.aws_iam_policy_document.external_dns_policy_doc.json
}

resource "kubernetes_service_account" "external_dns" {
  metadata {
    name      = "external-dns"
    namespace = "kube-system"
  }
}

resource "kubernetes_cluster_role" "external_dns" {
  metadata {
    name = "external-dns"
  }
  rule {
    verbs      = ["get", "list", "watch"]
    api_groups = [""]
    resources  = ["pods", "services"]
  }
  rule {
    verbs      = ["get", "list", "watch"]
    api_groups = ["extensions"]
    resources  = ["ingresses"]
  }
rule {
    verbs      = ["list"]
    api_groups = [""]
    resources  = ["nodes"]
  }
}

resource "kubernetes_cluster_role_binding" "external_dns" {
  metadata {
    name = "external-dns"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "external-dns"
    namespace = "kube-system"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "external-dns"
  }
}

resource "kubernetes_deployment" "external_dns" {
  metadata {
    name      = "external-dns"
    namespace = "kube-system"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app"    = "external-dns",
      }
    }
    template {
      metadata {
        labels = {
          "app"    = "external-dns",
        }
        annotations = {
          "iam.amazonaws.com/role" = "k8s-external-dns-role",
        }
      }
      spec {
        container {
          name  = "external-dns"
          image = "registry.opensource.zalan.do/teapot/external-dns:latest"
          args = [
            "--source=service",
            "--source=ingress",
            "--domain-filter=${var.external_dns_zone}", # Give access only to the specified zone
            "--provider=aws",
            "--aws-zone-type=public",
            "--policy=upsert-only", # Prevent ExternalDNS from deleting any records
            "--registry=txt",
            "--txt-owner-id=${var.external_dns_owner_id}", # ID of txt record to manage state
          ]
        }

        service_account_name = "external-dns"
        automount_service_account_token  = true
      }
    }
  }
}
