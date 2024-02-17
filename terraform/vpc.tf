resource "aws_vpc" "farmec_vpc" {
  cidr_block = "10.0.0.0/16"
  enable_dns_support = true
  enable_dns_hostnames = true
  tags = {
    Name = "farmec_vpc"
  }
}

resource "aws_subnet" "farmec_subnet" {
  vpc_id            = aws_vpc.farmec_vpc.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "eu-west-1a"
  map_public_ip_on_launch = true
  tags = {
    Name = "farmec_subnet"
  }
}

resource "aws_subnet" "farmec_subnet_2" {
  vpc_id            = aws_vpc.farmec_vpc.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "eu-west-1b"
  map_public_ip_on_launch = true
  tags = {
    Name = "farmec_subnet_2"
  }
}

resource "aws_subnet" "farmec_subnet_3" {
  vpc_id            = aws_vpc.farmec_vpc.id
  cidr_block        = "10.0.3.0/24"
  availability_zone = "eu-west-1c"
  map_public_ip_on_launch = true
  tags = {
    Name = "farmec_subnet_3"
  }
}

resource "aws_internet_gateway" "farmec_igw" {
  vpc_id = aws_vpc.farmec_vpc.id
  tags = {
    Name = "farmec_igw"
  }
}

resource "aws_route_table" "farmec_route_table" {
  vpc_id = aws_vpc.farmec_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.farmec_igw.id
  }

  tags = {
    Name = "farmec_route_table"
  }
}

resource "aws_route_table_association" "a" {
  subnet_id      = aws_subnet.farmec_subnet.id
  route_table_id = aws_route_table.farmec_route_table.id
}

resource "aws_route_table_association" "b" {
  subnet_id      = aws_subnet.farmec_subnet_2.id
  route_table_id = aws_route_table.farmec_route_table.id
}

resource "aws_route_table_association" "c" {
  subnet_id      = aws_subnet.farmec_subnet_3.id
  route_table_id = aws_route_table.farmec_route_table.id
}

