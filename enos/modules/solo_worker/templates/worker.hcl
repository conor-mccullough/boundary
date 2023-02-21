listener "tcp" {
  purpose = "proxy"
  tls_disable = true
  address = "0.0.0.0"
}

worker {
  # Name attr must be unique across workers
  name = "worker-${id}"
  description = "Enos Boundary worker ${id}"

  # Workers must be able to reach controllers on :9201
  controllers = ${controller_ips}

  public_addr = "${public_addr}"

  tags {
    type   = ["prod", "webservers"]
    region = ["us-east-1"]
  }
}

# must be same key as used on controller config
kms "awskms" {
  purpose    = "worker-auth"
  region     = "us-east-1"
  kms_key_id = "${kms_key_id}"
}
