output "instance_public_ip" {
  description = "Public IP of the EC2 Instance"
  value       = aws_instance.go_board_ec2.public_ip
}