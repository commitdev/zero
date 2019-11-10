terraform {
  required_version = ">= 0.12"
  backend "s3" {
    bucket         = "project-{{ .Config.Name }}-terraform-state"
    key            = "infrastructure/terraform/environments/production/main"
    encrypt        = true
    region         = "{{ .Config.Infrastructure.AWS.Region }}"
    dynamodb_table = "{{ .Config.Name }}-terraform-state-locks"
  }
}

# Instantiate the production environment
module "production" {
  source      = "../../modules/environment"
  environment = "production"

  # Project configuration
  project             = "{{ .Config.Infrastructure.AWS.EKS.ClusterName }}"
  region              = "{{ .Config.Infrastructure.AWS.Region }}"
  allowed_account_ids = ["{{ .Config.Infrastructure.AWS.AccountId }}"]

  {{- if ne .Config.Infrastructure.AWS.EKS.ClusterName "" }}
  # ECR configuration
  ecr_repositories = ["{{ .Config.Infrastructure.AWS.EKS.ClusterName }}"]

  # EKS configuration
  eks_worker_instance_type = "m4.large"
  eks_worker_asg_max_size  = 3

  # EKS-Optimized AMI for your region: https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html
  # https://us-east-1.console.aws.amazon.com/systems-manager/parameters/%252Faws%252Fservice%252Feks%252Foptimized-ami%252F1.14%252Famazon-linux-2%252Frecommended%252Fimage_id/description?region=us-east-1
  eks_worker_ami = "{{ .Config.Infrastructure.AWS.EKS.WorkerAMI }}"
  {{- end }}
}
