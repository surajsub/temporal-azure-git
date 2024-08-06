variable "location" {
  description = "Location to create this kubernetes resource in"
  
}
variable "rg_name" {
  description = "Name of the RG"

}

variable "aks_name" {
  description = "Name of the AKS Cluster"
}

variable "env" {
  description = "The environment value for it"
}

variable "rg_id" {
  description = "The id of the resource group"
}
variable "aks_subnet_id" {
  description = "The id of the subnet that is created for Azure K S"
}

variable "node_count" {
  description = "Node count for the AKS"
  default     = 1
}

variable "user_aid_id" {
  description = "The managed id that was created"
}
