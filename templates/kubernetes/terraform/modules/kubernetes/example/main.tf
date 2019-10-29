## Commit (example) service
## NOT CURRENTLY IMPORTED / INSTANTIATED ANYWHERE

# Create a role and establish a trust relationship with the worker nodes
resource "aws_iam_role" "k8s_worker_commit_service_role" {
  name                  = "k8s-service-commit"
  assume_role_policy    = data.aws_iam_policy_document.k8s_worker_assumerole_policy.json
  force_detach_policies = true
}

# Policy allowing access to specific AWS resources
data "aws_iam_policy_document" "k8s_commit_service_access_policy" {
  statement {
    actions   = ["s3:ListBucket"]
    resources = ["arn:aws:s3:::processed-data-files"]
  }

  statement {
    actions   = ["s3:GetObject"]
    resources = ["arn:aws:s3:::processed-data-files/*"]
  }
}

# Add the above policy to the created role
resource "aws_iam_role_policy" "k8s_commit_service_role_policy" {
  name   = "worker-commit-service-policy"
  role   = aws_iam_role.k8s_worker_commit_service_role.id
  policy = data.aws_iam_policy_document.k8s_commit_service_access_policy.json
}

