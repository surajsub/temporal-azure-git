resource "azurerm_user_assigned_identity" "temporal" {
  name                = "aks-temporal-identity"
  resource_group_name = var.rgname
  location            = var.location
}

resource "azurerm_role_assignment" "network_contributor" {
  scope                = var.rg_id
  role_definition_name = "Network Contributor"
  principal_id         = azurerm_user_assigned_identity.temporal.principal_id
}


output "mi_id" {
  value = azurerm_user_assigned_identity.temporal.id
}

output "client_id" {
  value = azurerm_user_assigned_identity.temporal.client_id
}

output "principal_id" {
  value = azurerm_user_assigned_identity.temporal.principal_id
}

output "tenant_id" {
  value = azurerm_user_assigned_identity.temporal.tenant_id
}