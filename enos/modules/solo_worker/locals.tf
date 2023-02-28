locals {
  name_prefix          = "${var.project_name}-${var.environment}"
  boundary_cluster_tag = "boundary-server-${random_string.cluster_id.result}"
}
