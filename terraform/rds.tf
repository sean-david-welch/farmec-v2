resource "aws_db_subnet_group" "farmec_db_subnet_group" {
  name       = "farmec_db_subnet_group"
  subnet_ids = [aws_subnet.farmec_subnet.id, aws_subnet.farmec_subnet_b.id, aws_subnet.farmec_subnet_c.id]  

  tags = {
    Name = "My DB Subnet Group"
  }
}

resource "aws_db_instance" "farmec_db_instance" {
  allocated_storage    = 20
  engine               = "postgres"
  engine_version       = "16" 
  instance_class       = "db.t3.micro"  
  publicly_accessible  = true
  username             = var.master_username
  password             = var.master_password
  parameter_group_name = "default.postgres16" 
  db_subnet_group_name = aws_db_subnet_group.farmec_db_subnet_group.name
  vpc_security_group_ids = [aws_security_group.my_sg.id]
  skip_final_snapshot  = true
}
