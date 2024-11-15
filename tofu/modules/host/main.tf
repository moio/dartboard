module "host" {
  source                = "../${var.backend}_host"
  project_name          = var.project_name
  name                  = var.name
  backend_variables = var.backend_variables
  ssh_key_name          = var.ssh_key_name
  ssh_private_key_path  = var.ssh_private_key_path
  ssh_user              = var.ssh_user
  ssh_bastion_host      = var.ssh_bastion_host
  ssh_bastion_user      = var.ssh_bastion_user
  ssh_tunnels = var.ssh_tunnels
  host_configuration_commands = var.host_configuration_commands
}
