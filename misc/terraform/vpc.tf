resource "aws_vpc" "farmec_vpc" {
  cidr_block = "10.0.0.0/16"  
  enable_dns_support = true
  enable_dns_hostnames = true

  tags = {
    Name = "FarmecVPC"
  }
}

resource "aws_subnet" "farmec_subnet" {
  vpc_id            = aws_vpc.farmec_vpc.id
  cidr_block        = "10.0.1.0/24"  
  availability_zone = "eu-west-1a"  
  map_public_ip_on_launch = true


  tags = {
    Name = "Farmec-Subnet-A"
  }
}

resource "aws_subnet" "farmec_subnet_b" {
  vpc_id            = aws_vpc.farmec_vpc.id
  cidr_block        = "10.0.2.0/24"  
  availability_zone = "eu-west-1b"  

  tags = {
    Name = "Farmec-Subnet-B"
  }
}

resource "aws_subnet" "farmec_subnet_c" {
  vpc_id            = aws_vpc.farmec_vpc.id
  cidr_block        = "10.0.3.0/24"  
  availability_zone = "eu-west-1c"  

  tags = {
    Name = "Farmec-Subnet-C"
  }
}

