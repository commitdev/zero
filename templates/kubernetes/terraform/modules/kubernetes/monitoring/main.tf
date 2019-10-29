resource "aws_iam_role" "k8s_monitoring" {
  name                  = "k8s-${var.environment}-monitoring"
  assume_role_policy    = var.assume_role_policy
  force_detach_policies = true
}

# Create amazon-cloudwatch kubernetes namespace for fluentd/cloudwatchagent
resource "kubernetes_namespace" "amazon_cloudwatch" {
  metadata {
    name = "amazon-cloudwatch"
    labels = {
      name = "amazon-cloudwatch"
    }
  }
}

data "aws_iam_policy" "CloudWatchAgentServerPolicy" {
  arn = "arn:aws:iam::aws:policy/CloudWatchAgentServerPolicy"
}

resource "aws_iam_role_policy_attachment" "k8s_monitoring_role_policy" {
  role       = "${aws_iam_role.k8s_monitoring.id}"
  policy_arn = "${data.aws_iam_policy.CloudWatchAgentServerPolicy.arn}"
}

module "fluentd" {
  source       = "./fluentd"
  environment  = var.environment
  region       = var.region
  cluster_name = var.cluster_name
}

module "cloudwatch_agent" {
  source       = "./cloudwatch_agent"
  environment  = var.environment
  region       = var.region
  cluster_name = var.cluster_name
}