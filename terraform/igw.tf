resource "aws_internet_gateway" "farmec_igw" {
  vpc_id = aws_vpc.farmec_vpc.id

  tags = {
    Name = "FarmecIGW"
  }
}

resource "aws_route_table" "farmec_route_table" {
  vpc_id = aws_vpc.farmec_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.farmec_igw.id
  }

  tags = {
    Name = "FarmecRouteTable"
  }
}

resource "aws_route_table_association" "a" {
  subnet_id      = aws_subnet.farmec_subnet.id
  route_table_id = aws_route_table.farmec_route_table.id
}

resource "aws_route_table_association" "b" {
  subnet_id      = aws_subnet.farmec_subnet_b.id
  route_table_id = aws_route_table.farmec_route_table.id
}

resource "aws_route_table_association" "c" {
  subnet_id      = aws_subnet.farmec_subnet_c.id
  route_table_id = aws_route_table.farmec_route_table.id
}
