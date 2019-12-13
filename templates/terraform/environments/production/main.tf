terraform {
  required_version = ">= 0.12"
  backend "s3" {
    bucket         = "{{ .Config.Name }}-production-terraform-state"
    key            = "infrastructure/terraform/environments/production/main"
    encrypt        = true
    region         = "{{ .Config.Infrastructure.AWS.Region }}"
    dynamodb_table = "{{ .Config.Name }}-production-terraform-state-locks"
  }
}

# Instantiate the production environment
module "production" {
  source      = "../../modules/environment"
  environment = "production"

  # Project configuration
  project             = "{{ .Config.Name }}"
  region              = "{{ .Config.Infrastructure.AWS.Region }}"
  allowed_account_ids = ["{{ .Config.Infrastructure.AWS.AccountId }}"]

  {{- if ne .Config.Infrastructure.AWS.EKS.ClusterName "" }}
  # ECR configuration
  ecr_repositories = ["{{ .Config.Infrastructure.AWS.EKS.ClusterName }}"]

  # EKS configuration
  eks_worker_instance_type = "m4.large"
  eks_worker_asg_min_size  = 3
  eks_worker_asg_max_size  = 6

  # EKS-Optimized AMI for your region: https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html
  # https://us-east-1.console.aws.amazon.com/systems-manager/parameters/%252Faws%252Fservice%252Feks%252Foptimized-ami%252F1.14%252Famazon-linux-2%252Frecommended%252Fimage_id/description?region=us-east-1
  eks_worker_ami = "{{ .Config.Infrastructure.AWS.EKS.WorkerAMI }}"
  {{- end }}

  {{- if .Config.Infrastructure.AWS.Cognito.Enabled }}
  # Cognito configuration
  user_pool = "{{ .Config.Name }}-production"
  hostname = "{{ .Config.Frontend.Hostname }}"
  {{- end}}

  # Hosting configuration
  s3_hosting_buckets = [
    "{{ .Config.Name }}-production"
  ]
  s3_hosting_cert_domain = "{{ .Config.Frontend.Hostname}}"

}

{{- if .Config.Infrastructure.AWS.Cognito.Enabled }}
output "cognito_client_id" {
  value = module.staging.cognito.cognito_client_id
}

output "cognito_pool_id" {
  value = module.staging.cognito.cognito_pool_id
}
{{- end}}
