resource "azurerm_public_ip" "temporal_ingress_pip" {
  name                = "temporal_ingress_pip"
  location            = var.rg_location
  resource_group_name = var.rg_name
  allocation_method   = "Static"
  sku                 = "Standard"

  tags = {
    environment = var.env
  }
}


output "public_ip" {
  value= azurerm_public_ip.temporal_ingress_pip.ip_address
}

output "public_ip_id" {
  value = azurerm_public_ip.temporal_ingress_pip.id
}
