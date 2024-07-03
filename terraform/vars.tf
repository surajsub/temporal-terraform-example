variable "region" {
  description = "The region to deploy this instance in"
  default =  "us-west-2"
}


variable "ami" {
  description = "AMI to deploy"
  default = "your-ami-id"
}

variable "subnet" {
  description = "The subnet to deploy this Instance in"
  default = "your-subnet-id"
}


variable "instance_type" {
  description = "The instance type"
  default = "t2.micro"
}
