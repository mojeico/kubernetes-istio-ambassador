### Backend and provider config for reference
terraform {
  required_version = "~> 0.12.12"
  backend "remote" {
    hostname = "app.terraform.io"
  }
}
provider "aws" {
  assume_role {
    role_arn     = var.assume_role_arn
    session_name = "EKS_deployment_session_${var.tags["Environment"]}"
  }
  version = "~> 2.28.1"
  region  = var.region
}

### EKS cluster config
resource "aws_eks_cluster" "cluster" {
  enabled_cluster_log_types = ["api", "audit", "authenticator", "controllerManager", "scheduler"]
  name                      = "${var.tags["ServiceType"]}-${var.tags["Environment"]}"
  role_arn                  = aws_iam_role.cluster.arn
  version                   = "1.14"
  vpc_config {
    subnet_ids              = flatten([module.vpc.public_subnets, module.vpc.private_subnets])
    security_group_ids      = [aws_security_group.cluster.id]
    endpoint_private_access = "true"
    endpoint_public_access  = "true"
  }
}
### OIDC config
resource "aws_iam_openid_connect_provider" "cluster" {
  client_id_list  = ["sts.amazonaws.com"]
  thumbprint_list = []
  url             = aws_eks_cluster.cluster.identity.0.oidc.0.issuer
}

### External cli kubergrunt
data "external" "thumb" {
  program = ["kubergrunt", "eks", "oidc-thumbprint", "--issuer-url", aws_eks_cluster.cluster.identity.0.oidc.0.issuer]
}

### OIDC config
resource "aws_iam_openid_connect_provider" "cluster" {
  client_id_list  = ["sts.amazonaws.com"]
  thumbprint_list = [data.external.thumb.result.thumbprint]
  url             = aws_eks_cluster.cluster.identity.0.oidc.0.issuer
}


data "aws_region" "current" {}
# Fetch OIDC provider thumbprint for root CA
data "external" "thumbprint" {
  program = ["./oidc-thumbprint.sh", data.aws_region.current.name]
}


resource "aws_iam_openid_connect_provider" "cluster" {
  client_id_list  = ["sts.amazonaws.com"]
  thumbprint_list = concat([data.external.thumbprint.result.thumbprint], var.oidc_thumbprint_list)
  url             = aws_eks_cluster.cluster.identity.0.oidc.0.issuer
}

