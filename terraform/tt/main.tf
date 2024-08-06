data "azurerm_kubernetes_cluster" "main_aks" {
  name                = "k8s1"
  resource_group_name = "temporal"
}
