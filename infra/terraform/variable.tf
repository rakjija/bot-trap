variable "aws_region" {
  description = "AWS region to deploy resources into"
  type        = string
}

variable "key_pair_name" {
  description = "Name of the existing EC2 key pair to use for SSH access"
  type        = string
}