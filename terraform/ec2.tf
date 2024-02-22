resource "aws_instance" "FarmecAPI" {
  ami           = "ami-0905a3c97561e0b69"
  instance_type = "t3.nano"
  subnet_id     = aws_subnet.farmec_subnet.id  
  key_name = aws_key_pair.farmec.key_name
  vpc_security_group_ids = [aws_security_group.my_ec2_sg.id]  

  tags = {
    Name = "FarmecAPI"
  }
}

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

resource "aws_route_table_association" "farmec_a_rta" {
  subnet_id      = aws_subnet.farmec_subnet.id
  route_table_id = aws_route_table.farmec_route_table.id
}
