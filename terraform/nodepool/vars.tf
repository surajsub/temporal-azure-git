variable "node_pool_name" {
  description = "Name of the node pool"
  default = "temporal"
}

variable "env" {
  description = "The environment to create the node pool in"
}


variable "vm_size" {
  description = "The vm size for this nodepool"
}


variable "aks_subnet_id" {
  description = "The id of the subnet that is created for Azure K S"
}


variable "aks_id" {
  description = "The kubernetes cluster id that was created"

}

variable "aks_version" {
  description = "The version of AKS"
}

