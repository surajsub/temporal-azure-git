resource "azurerm_resource_group" "temporal" {
  name     = var.name
  location = var.region
}


output "rg_id" {
  value = azurerm_resource_group.temporal.id
}

output "rg_name" {
  value = azurerm_resource_group.temporal.name
}