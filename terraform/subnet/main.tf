resource "azurerm_subnet" "aks_subnet" {
  name                 = "aks-subnet"
  address_prefixes     = ["10.8.0.0/21"]
  resource_group_name  = var.rg_name
  virtual_network_name = var.vnet_name
}

resource "azurerm_subnet" "applications_subnet" {
  name                 = "application-subnet"
  resource_group_name  = var.rg_name
  virtual_network_name =var.vnet_name
  address_prefixes     = ["10.8.8.0/21"]
}


output "aks_subnet_id" {
  value = azurerm_subnet.aks_subnet.id
}

output "app_subnet_id" {
  value= azurerm_subnet.applications_subnet.id
}

output "subnet_name" {
  value = azurerm_subnet.aks_subnet.name
}