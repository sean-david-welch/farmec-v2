resource "aws_instance" "FarmecAPI" {
  ami           = "ami-0905a3c97561e0b69"
  instance_type = "t3.nano"
  subnet_id     = aws_subnet.farmec_subnet.id
  key_name = aws_key_pair.deployer.key_name

  tags = {
    Name = "FarmecAPI"
  }
}