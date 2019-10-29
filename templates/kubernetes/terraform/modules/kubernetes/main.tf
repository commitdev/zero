module "kube2iam" {
  source      = "./kube2iam"
  environment = var.environment
  region      = var.region
}

module "monitoring" {
  source             = "./monitoring"
  environment        = var.environment
  region             = var.region
  assume_role_policy = var.assume_role_policy
  cluster_name       = var.cluster_name
}

module "ingress" {
  source                     = "./ingress"
  environment                = var.environment
  region                     = var.region
  load_balancer_ssl_cert_arn = ""
}