#
# Provider Configuration

provider "aws" {
  region = "${var.aws-region}"
}

resource "aws_key_pair" "tk8eks-key" {
  key_name = "tk8eks-key"
  # for tk8eks to allow ssh into the nodes, please adapt the path to your pub key below
  public_key = "${file(var.public_key_path)}"
}


# Using these data sources allows the configuration to be
# generic for any region.
data "aws_region" "current" {}

data "aws_availability_zones" "available" {}

provider "http" {}
