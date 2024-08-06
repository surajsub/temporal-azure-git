
variable "rg_name" {
  description = "The rg for this vnet to be created in"
}

variable "env" {
  description = "Environment that we are using"

}

variable "region" {
  description = "Region we want to deploy the resources in"
}


variable "vnet_block" {
  description = "The size of this vnet block"

}

variable "vnet_name" {
  description = "the name for this vnet"

}