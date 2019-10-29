resource "aws_ecr_repository" "ecr_repository" {
  count = length(var.ecr_repositories)
  name  = element(var.ecr_repositories, count.index)

  tags = {
    environment = var.environment
  }
}

