resource "azurerm_dns_zone" "temporal" {
  name                = var.dns_name
  resource_group_name = var.rg_name
}

resource "azurerm_dns_a_record" "ingress" {
  name                = "temporal"
  zone_name           = azurerm_dns_zone.temporal.name
  resource_group_name = var.rg_name
  ttl                 = 3600
  records             = [var.temporal_pip]
}



output "dns_id" {
  value = azurerm_dns_zone.temporal.id
}

output "dns_a_id" {
  value = azurerm_dns_a_record.ingress.id
}

output "dns_a_fqdn" {
  value = azurerm_dns_a_record.ingress.fqdn
}