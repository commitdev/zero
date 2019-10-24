# EKS Terraform 

AWS Resources created: 

- EKS Cluster: AWS managed Kubernetes cluster of master servers
- AutoScaling Group containing 2 m4.large instances based on the latest EKS Amazon Linux 2 AMI: Operator managed Kubernetes worker nodes for running Kubernetes service deployments
- Associated VPC, Internet Gateway, Security Groups, and Subnets: Operator managed networking resources for the EKS Cluster and worker node instances
- Associated IAM Roles and Policies: Operator managed access resources for EKS and worker node instances

## Pre-requisites

- Setup the [AWS credentials](https://www.terraform.io/docs/providers/aws/index.html#environment-variables) for terraform 

## Spin up cluster

```shell

terraform plan 
terraform apply 

```

### Connect to cluster
The EKS service does not provide a cluster-level API parameter or resource to automatically configure the underlying Kubernetes cluster to allow worker nodes to join the cluster via AWS IAM role authentication.

- Run `aws eks update-kubeconfig --name staging` to configure `kubectl`
- Run `terraform output config_map_aws_auth` and save the configuration into a file, e.g. config_map_aws_auth.yaml
- Run `kubectl apply -f config_map_aws_auth.yaml`
- You can verify the worker nodes are joining the cluster via: `kubectl get nodes --watch`
