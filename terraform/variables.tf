variable "aws_region" {
    description = "AWS region for all resources."

    type = string
    default = "eu-west-1"
}

variable "master_username" {
  default = ""
}

variable "master_password" {
  default = ""
}
