
#output "ingress_nginx_pip" {
#  value = data.azurerm_public_ip.ingress_nginx_pip.ip_address
#}


output "client_certificate" {
  value     = data.azurerm_kubernetes_cluster.main_aks.kube_config[0].client_certificate
  sensitive = true
}

output "client_key" {
  value     = data.azurerm_kubernetes_cluster.main_aks.kube_config[0].client_key
  sensitive = true
}

output "cluster_ca_certificate" {
  value     = data.azurerm_kubernetes_cluster.main_aks.kube_config[0].cluster_ca_certificate
  sensitive = true
}

output "cluster_password" {
  value     = data.azurerm_kubernetes_cluster.main_aks.kube_config[0].password
  sensitive = true
}

output "cluster_username" {
  value     = data.azurerm_kubernetes_cluster.main_aks.kube_config[0].username
  sensitive = true
}

output "host" {
  value     = data.azurerm_kubernetes_cluster.main_aks.kube_config[0].host
  sensitive = true
}

output "kube_config" {
  value     = data.azurerm_kubernetes_cluster.main_aks.kube_config_raw
  sensitive = true
}

