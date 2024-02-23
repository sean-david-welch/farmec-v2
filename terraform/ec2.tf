resource "aws_instance" "FarmecAPI" {
  ami           = "ami-0905a3c97561e0b69"
  instance_type = "t3.nano"
  subnet_id     = aws_subnet.farmec_subnet.id  
  key_name = aws_key_pair.farmec-ec2.key_name
  vpc_security_group_ids = [aws_security_group.my_ec2_sg.id]  

  tags = {
    Name = "FarmecAPI"
  }
}
