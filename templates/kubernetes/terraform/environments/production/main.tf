# Instantiate the production environment
module "production" {
  source      = "../../modules/environment"
  environment = "production"

  # Project configuration
  project             = "{{ .Kubernetes.ClusterName }}"
  region              = "{{ .Kubernetes.AWSRegion }}"
  allowed_account_ids = ["{{ .Kubernetes.AWSAccountId }}"]

  # ECR configuration
  ecr_repositories = ["{{ .Kubernetes.ClusterName }}"]

  # EKS configuration
  eks_worker_instance_type = "m4.large"
  eks_worker_asg_max_size  = 3

  # EKS-Optimized AMI for your region: https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html
  # https://us-east-1.console.aws.amazon.com/systems-manager/parameters/%252Faws%252Fservice%252Feks%252Foptimized-ami%252F1.14%252Famazon-linux-2%252Frecommended%252Fimage_id/description?region=us-east-1
  eks_worker_ami = "ami-0392bafc801b7520f"
}
