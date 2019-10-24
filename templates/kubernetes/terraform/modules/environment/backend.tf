terraform {
  backend "s3" {
    bucket         = "${var.project}-terraform-state"
    key            = "infrastructure/terraform/environments/${var.environment}/main"
    encrypt        = true
    region         = var.region
    dynamodb_table = "terraform-state-locks"
  }
}
