output "worker_ip" {
  description = "The public IP of the Boundary worker"
  value       = aws_instance.worker.public_ip
}
