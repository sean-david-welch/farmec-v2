provider "aws" {
  region = var.aws_region != "" ? var.aws_region : "${var.aws_region}"
}

resource "aws_key_pair" "farmec" {
  key_name   = "deployer-key"
  public_key = file("~/.ssh/farmec.pub")

  tags = {
    Name = "farmec"
  }
}

resource "aws_security_group" "my_sg" {
  name        = "my_sg"
  description = "Security Group for RDS cluster"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "my_sg"
  }
}

resource "aws_security_group" "my_alb_sg" {
  name        = "my_alb_sg"
  description = "Security Group for ALB"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "my_alb_sg"
  }
}
