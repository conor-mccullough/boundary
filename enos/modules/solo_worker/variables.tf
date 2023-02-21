variable "vpc_name" {
  description = "The name of the existing VPC to be used for this module"
  type        = string
}

variable "availability_zones" {
  description = "List of AWS availability zones to use (or * for all available)"
  type        = list(string)
  default     = ["*"]
}

variable "common_tags" {
  description = "Tags to set for all resources"
  type        = map(string)
  default     = { "Project" : "Enos" }
}

variable "kms_key_arn" {
  description = "ARN of KMS key used for SSHing to this module's instance"
  type        = string
}

variable "controller_ips" {
  description = "A list of public IPs for this worker's Boundary controllers"
  type        = list(string)
}

variable "ubuntu_ami_id" {
  description = "Ubuntu LTS AMI from the VPC this module will use"
  type        = string
}

variable "worker_instance_type" {
  description = "The EC2 Instance type to be used for the worker's node"
  type        = string
  default     = "t2.micro"
}

variable "ssh_aws_keypair" {
  description = "The name of the SSH keypair used to connect to EC2 instances"
  type        = string
}

variable "aws_iam_instance_profile_name" {
  description = "The name of the existing AWS instance profile to be used for this module"
  type        = string
}

variable "worker_monitoring" {
  description = "Enable detailed monitoring for workers"
  type        = bool
  default     = false
}

variable "worker_ebs_iops" {
  description = "EBS IOPS for the root volume"
  type        = number
  default     = null
}

variable "worker_ebs_size" {
  description = "EBS volume size"
  type        = number
  default     = 8
}

variable "worker_ebs_type" {
  description = "EBS volume type"
  type        = string
  default     = "gp2"
}

variable "worker_ebs_throughput" {
  description = "EBS data throughput (MiB/s) (only for gp2)"
  default     = null
}

variable "boundary_artifactory_release" {
  description = "Boundary release, version, and edition to install from artifactory.hashicorp.engineering"
  type = object({
    username = string
    token    = string
    url      = string
    sha256   = string
  })
  default = null
}

variable "local_artifact_path" {
  description = "The path to a local boundary.zip"
  type        = string
  default     = null
}

variable "boundary_release" {
  description = "Boundary release, version, and edition to install from releases.hashicorp.com"
  type = object({
    version = string
  })
  default = null
}
