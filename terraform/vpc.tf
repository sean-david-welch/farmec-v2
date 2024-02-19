data "aws_vpc" "default" {
  default = true
}
data "aws_subnet" "default_a" {
  id = "subnet-05c7858a1d43d94a6"
}

data "aws_subnet" "default_b" {
  id = "subnet-07ce5cdca67a7e3d1"
}

data "aws_subnet" "default_c" {
  id = "subnet-0586c91b03fe0244a"
}


