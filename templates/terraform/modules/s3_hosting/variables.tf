variable "project" {
  description = "The name of the project, mostly for tagging"
}

variable "buckets" {
  description = "S3 hosting buckets"
  type = set(string)
}

variable "cert_domain" {
  description = "Domain of the ACM certificate to lookup for Cloudfront to use"
  type = string
}
