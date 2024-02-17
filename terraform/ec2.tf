resource "aws_instance" "FarmecAPI" {
  ami           = "ami-0905a3c97561e0b69"
  instance_type = "t3.nano"
  subnet_id     = aws_subnet.farmec_subnet.id

  tags = {
    Name = "FarmecAPI"
  }
}