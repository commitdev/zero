data "local_file" "containers" {
  filename = "${path.module}/files/containers.conf"
}

data "local_file" "fluent" {
  filename = "${path.module}/files/fluent.conf"
}

data "local_file" "host" {
  filename = "${path.module}/files/host.conf"
}

data "local_file" "systemd" {
  filename = "${path.module}/files/systemd.conf"
}

resource "kubernetes_config_map" "cluster_info" {
  metadata {
    name      = "cluster-info"
    namespace = "amazon-cloudwatch"
  }
  data = {
    "cluster.name" = var.cluster_name
    "logs.region"  = var.region
  }
  depends_on = [kubernetes_namespace.amazon_cloudwatch]
}

resource "kubernetes_service_account" "fluentd" {
  metadata {
    name      = "fluentd"
    namespace = "amazon-cloudwatch"
  }
  depends_on = [kubernetes_namespace.amazon_cloudwatch]
}

resource "kubernetes_cluster_role" "fluentd_role" {
  metadata {
    name = "fluentd-role"
  }
  rule {
    verbs      = ["get", "list", "watch"]
    api_groups = [""]
    resources  = ["namespaces", "pods", "pods/logs"]
  }
}

resource "kubernetes_cluster_role_binding" "fluentd_role_binding" {
  metadata {
    name = "fluentd-role-binding"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "fluentd"
    namespace = "amazon-cloudwatch"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "fluentd-role"
  }
  depends_on = [kubernetes_service_account.fluentd]
}

resource "kubernetes_config_map" "fluentd_config" {
  metadata {
    name      = "fluentd-config"
    namespace = "amazon-cloudwatch"
    labels    = { k8s-app = "fluentd-cloudwatch" }
  }
  data = {
    "containers.conf" = data.local_file.containers.content
    "fluent.conf"     = data.local_file.fluent.content
    "host.conf"       = data.local_file.host.content
    "systemd.conf"    = data.local_file.systemd.content
  }
  depends_on = [kubernetes_namespace.amazon_cloudwatch]
}

resource "kubernetes_daemonset" "fluentd_cloudwatch" {
  depends_on = [
    kubernetes_config_map.cluster_info,
    kubernetes_config_map.fluentd_config
  ]
  metadata {
    name      = "fluentd-cloudwatch"
    namespace = "amazon-cloudwatch"
    labels = {
      k8s-app = "fluentd-cloudwatch"
    }
  }
  spec {
    selector {
      match_labels = {
        k8s-app = "fluentd-cloudwatch"
      }
    }
    template {
      metadata {
        labels = {
          k8s-app = "fluentd-cloudwatch"
        }
        annotations = {
          configHash               = "8915de4cf9c3551a8dc74c0137a3e83569d28c71044b0359c2578d2e0461825",
          "iam.amazonaws.com/role" = "k8s-${var.environment}-monitoring"
        }
      }
      spec {
        volume {
          name = "config-volume"
          config_map {
            name = "fluentd-config"
          }
        }
        volume {
          name = "fluentdconf"
        }
        volume {
          name = "varlog"
          host_path {
            path = "/var/log"
          }
        }
        volume {
          name = "varlibdockercontainers"
          host_path {
            path = "/var/lib/docker/containers"
          }
        }
        volume {
          name = "runlogjournal"
          host_path {
            path = "/run/log/journal"
          }
        }
        volume {
          name = "dmesg"
          host_path {
            path = "/var/log/dmesg"
          }
        }
        init_container {
          name    = "copy-fluentd-config"
          image   = "busybox"
          command = ["sh", "-c", "cp /config-volume/..data/* /fluentd/etc"]
          volume_mount {
            name       = "config-volume"
            mount_path = "/config-volume"
          }
          volume_mount {
            name       = "fluentdconf"
            mount_path = "/fluentd/etc"
          }
        }
        init_container {
          name    = "update-log-driver"
          image   = "busybox"
          command = ["sh", "-c", ""]
        }
        container {
          name  = "fluentd-cloudwatch"
          image = "fluent/fluentd-kubernetes-daemonset:v1.3.3-debian-cloudwatch-1.4"
          env {
            name = "REGION"
            value_from {
              config_map_key_ref {
                name = "cluster-info"
                key  = "logs.region"
              }
            }
          }
          env {
            name = "CLUSTER_NAME"
            value_from {
              config_map_key_ref {
                name = "cluster-info"
                key  = "cluster.name"
              }
            }
          }
          resources {
            limits {
              memory = "200Mi"
            }
            requests {
              cpu    = "100m"
              memory = "200Mi"
            }
          }
          volume_mount {
            name       = "config-volume"
            mount_path = "/config-volume"
          }
          volume_mount {
            name       = "fluentdconf"
            mount_path = "/fluentd/etc"
          }
          volume_mount {
            name       = "varlog"
            mount_path = "/var/log"
          }
          volume_mount {
            name       = "varlibdockercontainers"
            read_only  = true
            mount_path = "/var/lib/docker/containers"
          }
          volume_mount {
            name       = "runlogjournal"
            read_only  = true
            mount_path = "/run/log/journal"
          }
          volume_mount {
            name       = "dmesg"
            read_only  = true
            mount_path = "/var/log/dmesg"
          }
        }
        termination_grace_period_seconds = 30
        service_account_name             = "fluentd"
        automount_service_account_token  = true
      }
    }
  }
}
