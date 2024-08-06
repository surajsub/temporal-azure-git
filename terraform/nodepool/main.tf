resource "azurerm_kubernetes_cluster_node_pool" "temporal" {
  name                  = var.node_pool_name
  kubernetes_cluster_id = var.aks_id
  vm_size               = var.vm_size
  node_count            = 1
  min_count             = 1
  max_count             = 5
  vnet_subnet_id        = var.aks_subnet_id
  orchestrator_version  = var.aks_version

  enable_auto_scaling = true

  node_labels = {
    role = "worker01"
  }

  tags = {
    Environment = var.env
  }

  lifecycle {
    ignore_changes = [node_count]
  }
}