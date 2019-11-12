resource "kubernetes_namespace" "ingress_nginx" {
  metadata {
    name = "ingress-nginx"
    labels = {
      "app.kubernetes.io/name"    = "ingress-nginx",
      "app.kubernetes.io/part-of" = "ingress-nginx"
    }
  }
}

resource "kubernetes_config_map" "nginx_configuration" {
  metadata {
    name      = "nginx-configuration"
    namespace = "ingress-nginx"
    labels = {
      "app.kubernetes.io/name"    = "ingress-nginx",
      "app.kubernetes.io/part-of" = "ingress-nginx"
    }
  }
  data = {
    proxy-real-ip-cidr    = "0.0.0.0/0",
    use-forwarded-headers = "true",
    use-proxy-protocol    = "false"
  }
  depends_on = [kubernetes_namespace.ingress_nginx]
}

# resource "kubernetes_config_map" "nginx_configuration" {
#   metadata {
#     name      = "nginx-configuration"
#     namespace = "ingress-nginx"
#     labels = {
#       "app.kubernetes.io/name"    = "ingress-nginx",
#       "app.kubernetes.io/part-of" = "ingress-nginx"
#     }
#   }
#  depends_on = [kubernetes_namespace.ingress_nginx]
# }

resource "kubernetes_config_map" "tcp_services" {
  metadata {
    name      = "tcp-services"
    namespace = "ingress-nginx"
    labels = {
      "app.kubernetes.io/name"    = "ingress-nginx",
      "app.kubernetes.io/part-of" = "ingress-nginx"
    }
  }
  depends_on = [kubernetes_namespace.ingress_nginx]
}

resource "kubernetes_config_map" "udp_services" {
  metadata {
    name      = "udp-services"
    namespace = "ingress-nginx"
    labels = {
      "app.kubernetes.io/name"    = "ingress-nginx",
      "app.kubernetes.io/part-of" = "ingress-nginx"
    }
  }
  depends_on = [kubernetes_namespace.ingress_nginx]
}

resource "kubernetes_service_account" "nginx_ingress_serviceaccount" {
  metadata {
    name      = "nginx-ingress-serviceaccount"
    namespace = "ingress-nginx"
    labels = {
      "app.kubernetes.io/name"    = "ingress-nginx",
      "app.kubernetes.io/part-of" = "ingress-nginx"
    }
  }
  depends_on = [kubernetes_namespace.ingress_nginx]
}

resource "kubernetes_cluster_role" "nginx_ingress_clusterrole" {
  metadata {
    name = "nginx-ingress-clusterrole"
    labels = {
      "app.kubernetes.io/name"    = "ingress-nginx",
      "app.kubernetes.io/part-of" = "ingress-nginx"
    }
  }
  rule {
    verbs      = ["list", "watch"]
    api_groups = [""]
    resources  = ["configmaps", "endpoints", "nodes", "pods", "secrets"]
  }
  rule {
    verbs      = ["get"]
    api_groups = [""]
    resources  = ["nodes"]
  }
  rule {
    verbs      = ["get", "list", "watch"]
    api_groups = [""]
    resources  = ["services"]
  }
  rule {
    verbs      = ["create", "patch"]
    api_groups = [""]
    resources  = ["events"]
  }
  rule {
    verbs      = ["get", "list", "watch"]
    api_groups = ["extensions", "networking.k8s.io"]
    resources  = ["ingresses"]
  }
  rule {
    verbs      = ["update"]
    api_groups = ["extensions", "networking.k8s.io"]
    resources  = ["ingresses/status"]
  }
}

resource "kubernetes_role" "nginx_ingress_role" {
  metadata {
    name      = "nginx-ingress-role"
    namespace = "ingress-nginx"
    labels = {
      "app.kubernetes.io/name"    = "ingress-nginx",
      "app.kubernetes.io/part-of" = "ingress-nginx"
    }
  }
  rule {
    verbs      = ["get"]
    api_groups = [""]
    resources  = ["configmaps", "pods", "secrets", "namespaces"]
  }
  rule {
    verbs          = ["get", "update"]
    api_groups     = [""]
    resources      = ["configmaps"]
    resource_names = ["ingress-controller-leader-nginx"]
  }
  rule {
    verbs      = ["create"]
    api_groups = [""]
    resources  = ["configmaps"]
  }
  rule {
    verbs      = ["get"]
    api_groups = [""]
    resources  = ["endpoints"]
  }
  depends_on = [kubernetes_namespace.ingress_nginx]
}

resource "kubernetes_role_binding" "nginx_ingress_role_nisa_binding" {
  metadata {
    name      = "nginx-ingress-role-nisa-binding"
    namespace = "ingress-nginx"
    labels = {
      "app.kubernetes.io/name"    = "ingress-nginx",
      "app.kubernetes.io/part-of" = "ingress-nginx"
    }
  }
  subject {
    kind      = "ServiceAccount"
    name      = "nginx-ingress-serviceaccount"
    namespace = "ingress-nginx"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "Role"
    name      = "nginx-ingress-role"
  }
  depends_on = [kubernetes_namespace.ingress_nginx]
}

resource "kubernetes_cluster_role_binding" "nginx_ingress_clusterrole_nisa_binding" {
  metadata {
    name = "nginx-ingress-clusterrole-nisa-binding"
    labels = {
      "app.kubernetes.io/name"    = "ingress-nginx",
      "app.kubernetes.io/part-of" = "ingress-nginx"
    }
  }
  subject {
    kind      = "ServiceAccount"
    name      = "nginx-ingress-serviceaccount"
    namespace = "ingress-nginx"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "nginx-ingress-clusterrole"
  }
}

resource "kubernetes_deployment" "nginx_ingress_controller" {
  depends_on = [
    kubernetes_config_map.tcp_services,
    kubernetes_config_map.udp_services,
  ]
  metadata {
    name      = "nginx-ingress-controller"
    namespace = "ingress-nginx"
    labels = {
      "app.kubernetes.io/name"    = "ingress-nginx",
      "app.kubernetes.io/part-of" = "ingress-nginx"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app.kubernetes.io/name"    = "ingress-nginx",
        "app.kubernetes.io/part-of" = "ingress-nginx"
      }
    }
    template {
      metadata {
        labels = {
          "app.kubernetes.io/name"    = "ingress-nginx",
          "app.kubernetes.io/part-of" = "ingress-nginx"
        }
        # annotations = {
        #   "prometheus.io/port"   = "10254",
        #   "prometheus.io/scrape" = "true"
        # }
      }
      spec {
        container {
          name  = "nginx-ingress-controller"
          image = "quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.26.1"
          args = [
            "/nginx-ingress-controller",
            "--configmap=$(POD_NAMESPACE)/nginx-configuration",
            "--tcp-services-configmap=$(POD_NAMESPACE)/tcp-services",
            "--udp-services-configmap=$(POD_NAMESPACE)/udp-services",
            "--publish-service=$(POD_NAMESPACE)/ingress-nginx",
            "--annotations-prefix=nginx.ingress.kubernetes.io"
          ]
          port {
            name           = "http"
            container_port = 80
          }
          port {
            name           = "https"
            container_port = 443
          }
          env {
            name = "POD_NAME"
            value_from {
              field_ref {
                field_path = "metadata.name"
              }
            }
          }
          env {
            name = "POD_NAMESPACE"
            value_from {
              field_ref {
                field_path = "metadata.namespace"
              }
            }
          }
          liveness_probe {
            http_get {
              path   = "/healthz"
              port   = "10254"
              scheme = "HTTP"
            }
            initial_delay_seconds = 10
            timeout_seconds       = 10
            period_seconds        = 10
            success_threshold     = 1
            failure_threshold     = 3
          }
          readiness_probe {
            http_get {
              path   = "/healthz"
              port   = "10254"
              scheme = "HTTP"
            }
            timeout_seconds   = 10
            period_seconds    = 10
            success_threshold = 1
            failure_threshold = 3
          }
          lifecycle {
            pre_stop {
              exec {
                command = ["/wait-shutdown"]
              }
            }
          }
          security_context {
            run_as_user                = 33
            allow_privilege_escalation = true
            capabilities {
              add  = ["NET_BIND_SERVICE"]
              drop = ["ALL"]
            }
          }
        }
        termination_grace_period_seconds = 300
        node_selector                    = { "kubernetes.io/os" = "linux" }
        service_account_name             = "nginx-ingress-serviceaccount"
        automount_service_account_token  = true
      }
    }
  }
}

resource "kubernetes_service" "ingress_nginx" {
  metadata {
    name      = "ingress-nginx"
    namespace = "ingress-nginx"
    labels = {
      "app.kubernetes.io/name"    = "ingress-nginx",
      "app.kubernetes.io/part-of" = "ingress-nginx"
    }
  }
  spec {
    port {
      name        = "http"
      port        = 80
      target_port = "http"
    }
    port {
      name        = "https"
      port        = 443
      target_port = "https"
    }
    selector = {
      "app.kubernetes.io/name"    = "ingress-nginx",
      "app.kubernetes.io/part-of" = "ingress-nginx"
    }
    type                    = "LoadBalancer"
    external_traffic_policy = "Local"
  }
  depends_on = [kubernetes_namespace.ingress_nginx]
}

# HTTPS Load balancer
# resource "kubernetes_service" "ingress_nginx" {
#   metadata {
#     name      = "ingress-nginx"
#     namespace = "ingress-nginx"
#     labels = {
#       "app.kubernetes.io/name"    = "ingress-nginx",
#       "app.kubernetes.io/part-of" = "ingress-nginx"
#     }
#     annotations = {
#       "service.beta.kubernetes.io/aws-load-balancer-backend-protocol"        = "http",
#       "service.beta.kubernetes.io/aws-load-balancer-connection-idle-timeout" = "60",
#       "service.beta.kubernetes.io/aws-load-balancer-ssl-cert"                = var.load_balancer_ssl_cert_arn,
#       "service.beta.kubernetes.io/aws-load-balancer-ssl-ports"               = "https"
#     }
#   }
#   spec {
#     port {
#       name        = "http"
#       port        = 80
#       target_port = "http"
#     }
#     port {
#       name        = "https"
#       port        = 443
#       target_port = "http"
#     }
#     selector = {
#       "app.kubernetes.io/name"    = "ingress-nginx",
#       "app.kubernetes.io/part-of" = "ingress-nginx"
#     }
#     type = "LoadBalancer"
#   }
#  depends_on = [kubernetes_namespace.ingress_nginx]
# }
