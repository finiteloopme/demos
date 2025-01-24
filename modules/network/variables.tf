variable "project_id" {
  description = "ID for the project"
}
variable "network_name" {
  description = "Name of the network"
  default     = "default"
}

variable "auto_create_subnetworks" {
  description = "value for auto_create_subnetworks"
  default     = true
}
