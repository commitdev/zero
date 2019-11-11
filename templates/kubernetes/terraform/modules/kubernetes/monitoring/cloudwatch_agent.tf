resource "kubernetes_service_account" "cloudwatch_agent" {
  metadata {
    name      = "cloudwatch-agent"
    namespace = "amazon-cloudwatch"
  }
  depends_on = [kubernetes_namespace.amazon_cloudwatch]
}

resource "kubernetes_cluster_role" "cloudwatch_agent_role" {
  metadata {
    name = "cloudwatch-agent-role"
  }
  rule {
    verbs      = ["list", "watch"]
    api_groups = [""]
    resources  = ["pods", "nodes", "endpoints"]
  }
  rule {
    verbs      = ["list", "watch"]
    api_groups = ["apps"]
    resources  = ["replicasets"]
  }
  rule {
    verbs      = ["list", "watch"]
    api_groups = ["batch"]
    resources  = ["jobs"]
  }
  rule {
    verbs      = ["get"]
    api_groups = [""]
    resources  = ["nodes/proxy"]
  }
  rule {
    verbs      = ["create"]
    api_groups = [""]
    resources  = ["nodes/stats", "configmaps", "events"]
  }
  rule {
    verbs          = ["get", "update"]
    api_groups     = [""]
    resources      = ["configmaps"]
    resource_names = ["cwagent-clusterleader"]
  }
}

resource "kubernetes_cluster_role_binding" "cloudwatch_agent_role_binding" {
  metadata {
    name = "cloudwatch-agent-role-binding"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "cloudwatch-agent"
    namespace = "amazon-cloudwatch"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "cloudwatch-agent-role"
  }
}

resource "kubernetes_config_map" "cwagentconfig" {
  metadata {
    name      = "cwagentconfig"
    namespace = "amazon-cloudwatch"
  }
  data = {
    "cwagentconfig.json" = templatefile(
      "${path.module}/files/cwagentconfig.json.tpl",
      {
        region       = var.region,
        cluster_name = var.cluster_name
      }
    )
  }
  depends_on = [kubernetes_namespace.amazon_cloudwatch]
}

resource "kubernetes_daemonset" "cloudwatch_agent" {
  # Explicitly declare dependency on config map
  depends_on = [
    kubernetes_config_map.cwagentconfig
  ]
  metadata {
    name      = "cloudwatch-agent"
    namespace = "amazon-cloudwatch"
  }
  spec {
    selector {
      match_labels = { name = "cloudwatch-agent" }
    }
    template {
      metadata {
        labels = { name = "cloudwatch-agent" }
        annotations = {
          "iam.amazonaws.com/role" = "k8s-${var.environment}-monitoring"
        }
      }
      spec {
        volume {
          name = "cwagentconfig"
          config_map {
            name = "cwagentconfig"
          }
        }
        volume {
          name = "rootfs"
          host_path {
            path = "/"
          }
        }
        volume {
          name = "dockersock"
          host_path {
            path = "/var/run/docker.sock"
          }
        }
        volume {
          name = "varlibdocker"
          host_path {
            path = "/var/lib/docker"
          }
        }
        volume {
          name = "sys"
          host_path {
            path = "/sys"
          }
        }
        volume {
          name = "devdisk"
          host_path {
            path = "/dev/disk/"
          }
        }
        container {
          name  = "cloudwatch-agent"
          image = "amazon/cloudwatch-agent:latest"
          port {
            container_port = 8125
            host_port      = 8125
            protocol       = "UDP"
          }
          env {
            name = "HOST_IP"
            value_from {
              field_ref {
                field_path = "status.hostIP"
              }
            }
          }
          env {
            name = "HOST_NAME"
            value_from {
              field_ref {
                field_path = "spec.nodeName"
              }
            }
          }
          env {
            name = "K8S_NAMESPACE"
            value_from {
              field_ref {
                field_path = "metadata.namespace"
              }
            }
          }
          env {
            name  = "CI_VERSION"
            value = "k8s/1.0.0"
          }
          resources {
            limits {
              cpu    = "200m"
              memory = "200Mi"
            }
            requests {
              memory = "200Mi"
              cpu    = "200m"
            }
          }
          volume_mount {
            name       = "cwagentconfig"
            mount_path = "/etc/cwagentconfig"
          }
          volume_mount {
            name       = "rootfs"
            read_only  = true
            mount_path = "/rootfs"
          }
          volume_mount {
            name       = "dockersock"
            read_only  = true
            mount_path = "/var/run/docker.sock"
          }
          volume_mount {
            name       = "varlibdocker"
            read_only  = true
            mount_path = "/var/lib/docker"
          }
          volume_mount {
            name       = "sys"
            read_only  = true
            mount_path = "/sys"
          }
          volume_mount {
            name       = "devdisk"
            read_only  = true
            mount_path = "/dev/disk"
          }
        }
        termination_grace_period_seconds = 60
        service_account_name             = "cloudwatch-agent"
        automount_service_account_token  = true
      }
    }
  }
}
