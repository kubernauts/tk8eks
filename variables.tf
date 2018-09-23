# Variables Configuration
variable "cluster-name" {
  default     = "eks1-kubernauts-eu-west-1"
  type        = "string"
  description = "The name of your EKS Cluster"
}

variable "aws-region" {
  default     = "eu-west-1"
  # availabe regions are:
  # us-east-1 (Virginia)
  # us-west-2 (Oregon)
  # eu-west-1 (Irland)
  type        = "string"
  description = "The AWS Region to deploy EKS"
}

variable "node-instance-type" {
  default     = "m4.large"
  type        = "string"
  description = "Worker Node EC2 instance type"
}

variable "desired-capacity" {
  default     = 2
  type        = "string"
  description = "Autoscaling Desired node capacity"
}

variable "max-size" {
  default     = 5
  type        = "string"
  description = "Autoscaling maximum node capacity"
}

variable "min-size" {
  default     = 1
  type        = "string"
  description = "Autoscaling Minimum node capacity"
}
