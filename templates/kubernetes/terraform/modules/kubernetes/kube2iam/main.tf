resource "kubernetes_service_account" "kube2iam" {
  metadata {
    name      = "kube2iam"
    namespace = "kube-system"
  }
}

resource "kubernetes_cluster_role" "kube2iam" {
  metadata {
    name = "kube2iam"
  }
  rule {
    verbs      = ["get", "watch", "list"]
    api_groups = [""]
    resources  = ["namespaces", "pods"]
  }
}

resource "kubernetes_cluster_role_binding" "kube2iam" {
  metadata {
    name = "kube2iam"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "kube2iam"
    namespace = "kube-system"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "kube2iam"
  }
}

resource "kubernetes_daemonset" "kube2iam" {
  metadata {
    name      = "kube2iam"
    namespace = "kube-system"
    labels = {
      app = "kube2iam"
    }
  }
  spec {
    selector {
      match_labels = {
        name = "kube2iam"
      }
    }
    template {
      metadata {
        labels = {
          name = "kube2iam"
        }
      }
      spec {
        container {
          name  = "kube2iam"
          image = "jtblin/kube2iam:0.10.8"
          args = [
            "--auto-discover-base-arn",
            "--auto-discover-default-role",
            "--iptables=true",
            "--host-ip=$(HOST_IP)",
            "--host-interface=eni+",
            # "--node=$(NODE_NAME)",
            "--use-regional-sts-endpoint",
            "--log-level=info"
          ]
          port {
            name           = "http"
            host_port      = 8181
            container_port = 8181
          }
          env {
            name = "HOST_IP"
            value_from {
              field_ref {
                field_path = "status.podIP"
              }
            }
          }
          # env {
          #   name = "NODE_NAME"
          #   value_from {
          #     field_ref {
          #       field_path = "spec.nodeName"
          #     }
          #   }
          # }
          env {
            name  = "AWS_REGION"
            value = var.region
          }
          security_context {
            privileged = true
          }
        }
        service_account_name            = "kube2iam"
        automount_service_account_token = true
        host_network                    = true
      }
    }
  }
}

