resource "aws_lb" "farmec_loadbalancer" {
  name               = "farmec-loadbalancer"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.my_alb_sg.id]
  subnets            = [aws_subnet.farmec_subnet.id, aws_subnet.farmec_subnet_2.id]

  enable_deletion_protection = false

  tags = {
    Name = "farmec-loadbalancer"
  }
}

resource "aws_lb_target_group" "my_tg" {
  name     = "my-tg"
  port     = 80
  protocol = "HTTP"
  vpc_id   = aws_vpc.farmec_vpc.id
}

resource "aws_lb_listener" "my_listener" {
  load_balancer_arn = aws_lb.farmec_loadbalancer.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.my_tg.arn
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
    cidr_blocks = ["0.0.0.0/0"]
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

resource "aws_security_group" "my_alb_sg" {
  name        = "my_alb_sg"
  description = "Security Group for ALB"
  vpc_id      = aws_vpc.farmec_vpc.id

  ingress {
    from_port   = 80
    to_port     = 80
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
    Name = "my_alb_sg"
  }
}