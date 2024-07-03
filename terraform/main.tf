
provider aws {
  region = var.region
}

resource "aws_instance" "temporal_server" {


  ami           = var.ami
  instance_type = var.instance_type
  subnet_id     = var.subnet



  tags = {

    Name = "suraj-temporal"

  }

}

output "instance_id" {
  value = aws_instance.temporal_server.id
}

output "instance_public_ip" {
  value = aws_instance.temporal_server.public_ip
}

