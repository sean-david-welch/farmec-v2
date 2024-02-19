resource "aws_instance" "FarmecAPI" {
  ami           = "ami-0905a3c97561e0b69"
  instance_type = "t3.nano"
  subnet_id     = tolist(data.aws_subnet_ids.default.ids)[0] 
  key_name = aws_key_pair.farmec.key_name

  tags = {
    Name = "FarmecAPI"
  }
}