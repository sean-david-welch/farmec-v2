provider "aws" {
  region = var.aws_region != "" ? var.aws_region : "${var.aws_region}"
}

resource "aws_key_pair" "farmec-ec2" {
  key_name   = "farmec-ec2"
  public_key = file("~/.ssh/farmec.pub")

  tags = {
    Name = "farmec-ec2"
  }
}

resource "aws_security_group" "my_sg" {
  name        = "my_sg"
  description = "Security Group for RDS cluster"
  vpc_id      = aws_vpc.farmec_vpc.id  

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = [
      "10.0.1.0/24",
      "10.0.2.0/24",
      "0.0.0.0/0"   
    ]
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

resource "aws_security_group" "my_ec2_sg" {
  name        = "my_ec2_sg"
  description = "Security Group for EC2 instances"
  vpc_id      = aws_vpc.farmec_vpc.id

  ingress {
    from_port   = 22
    to_port     = 22
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
    Name = "SSH Access SG"
  }
}


