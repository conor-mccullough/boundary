terraform {
  required_providers {
    enos = {
      source = "app.terraform.io/hashicorp-qti/enos"
    }
  }
}

resource "random_string" "cluster_id" {
  length  = 8
  upper   = false
  numeric = false
  special = false
}

resource "random_pet" "worker" {
  keepers = {
    # Generate a new pet name each time the instance has a new ARN
    ami_id = aws_instance.worker.arn
  }
  separator = "_"
}

data "aws_availability_zones" "available" {
  state = "available"
  filter {
    name   = "zone-name"
    values = var.availability_zones
  }
}

data "aws_kms_key" "kms_key" {
  key_id = var.kms_key_arn
}

resource "random_integer" "az" {
  min = 0
  max = length(data.aws_availability_zones.available.names) - 1
  keepers = {
    # Generate a new integer each time the list of aws_availability_zones changes
    # keepers have to be strings, sort the list in case order changes but zones don't
    listener_arn = join("", sort(data.aws_availability_zones.available.names))
  }
}

resource "aws_subnet" "default" {
  vpc_id                  = var.vpc_name
  cidr_block              = "10.13.10.0/24"
  map_public_ip_on_launch = true
  availability_zone       = data.aws_availability_zones.available.names[random_integer.az.result]
  tags = merge(
    var.common_tags,
    {
      "Name" = "${var.vpc_name}_solo_worker_subnet"
    },
  )
}

resource "aws_security_group" "default" {
  vpc_id = var.vpc_name

  ingress {
    description = "allow traffic from all IPs"
    from_port   = 0
    to_port     = 0
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = merge(
    var.common_tags,
    {
      "Name" = "${var.vpc_name}_solo_worker_sg"
    },
  )
}

resource "aws_instance" "worker" {
  ami                    = var.ubuntu_ami_id
  instance_type          = var.worker_instance_type
  vpc_security_group_ids = [aws_security_group.default.id]
  subnet_id              = aws_subnet.default.id
  key_name               = var.ssh_aws_keypair
  iam_instance_profile   = var.iam_instance_profile_name
  monitoring             = var.worker_monitoring

  root_block_device {
    iops        = var.ebs_iops
    volume_size = var.ebs_size
    volume_type = var.ebs_type
    throughput  = var.ebs_throughput
    tags        = var.common_tags
  }

  tags = merge(
    var.common_tags,
    {
      Name = "${var.name_prefix}-boundary-solo-worker",
      Type = var.cluster_tag,
    },
  )
}

resource "enos_bundle_install" "worker" {
  depends_on  = [aws_instance.worker]
  destination = var.boundary_install_dir

  artifactory = var.boundary_artifactory_release
  path        = var.local_artifact_path
  release     = var.boundary_release == null ? var.boundary_release : merge(var.boundary_release, { product = "boundary", edition = "oss" })

  transport = {
    ssh = {
      host = aws_instance.worker.public_ip
    }
  }
}

resource "enos_file" "worker_config" {
  depends_on  = [enos_bundle_install.worker]
  destination = "/etc/boundary/boundary.hcl"
  content = templatefile("${path.module}/templates/worker.hcl", {
    id             = random_pet.worker.id
    kms_key_id     = data.aws_kms_key.kms_key.id,
    controller_ips = jsonencode(var.controller_ips),
    public_addr    = aws_instance.worker.public_ip
  })

  transport = {
    ssh = {
      host = aws_instance.worker.public_ip
    }
  }
}

resource "enos_boundary_start" "worker_start" {
  depends_on = [enos_file.worker_config]

  bin_path    = "/opt/boundary/bin"
  config_path = "/etc/boundary"
  transport = {
    ssh = {
      host = aws_instance.worker.public_ip
    }
  }
}
