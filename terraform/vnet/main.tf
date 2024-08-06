resource "azurerm_virtual_network" "temporal" {
  name                = var.vnet_name
  address_space       = [var.vnet_block]
  location            = var.region
  resource_group_name = var.rg_name

  tags = {
    env = var.env
  }
}



output "vnet_id" {
  value = azurerm_virtual_network.temporal.id
}
output "vnet_guid_id" {
  value = azurerm_virtual_network.temporal.guid
}

output "vnet_name" {
  value  = azurerm_virtual_network.temporal.name
}