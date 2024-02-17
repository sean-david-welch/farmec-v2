provider "aws" {
  region = var.aws_region != "" ? var.aws_region : "${var.aws_region}"
}





