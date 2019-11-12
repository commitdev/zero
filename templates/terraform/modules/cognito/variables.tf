variable "user_pool" {
  description = "AWS Cognito pool name"
} 
variable "hostname" {
  default = "{{ .Config.Frontend.Hostname }}"
  description = "AWS Cognito pool name"
} 
