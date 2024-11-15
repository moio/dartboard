variable "project_name" {
  description = "A prefix for names of objects created by this module"
  default     = "st"
}

variable "name" {
  description = "Symbolic name of this instance"
  type        = string
}

variable "backend" {
  description = "Backend for this host"
  type        = string
}

variable "backend_variables" {
  description = "Backend-specific configuration variables"
  type = any
}

variable "ssh_key_name" {
  description = "Name of the SSH key used to access the instance"
  type        = string
}

variable "ssh_private_key_path" {
  description = "Path of private ssh key used to access the instance"
  type        = string
}

variable "ssh_user" {
  description = "User name to use for the SSH connection"
  type        = string
  default     = "root"
}

variable "ssh_bastion_host" {
  description = "Public name of the SSH bastion host. Leave null for publicly accessible instances"
  type        = string
  default     = null
}

variable "ssh_bastion_user" {
  description = "User name for the SSH bastion host's OS"
  default     = "root"
}

variable "ssh_tunnels" {
  description = "Opens SSH tunnels to this host via the bastion"
  type        = list(list(number))
  default     = []
}

variable "host_configuration_commands" {
  description = "Commands to run when the host is deployed"
  default     = ["cat /etc/os-release"]
}
