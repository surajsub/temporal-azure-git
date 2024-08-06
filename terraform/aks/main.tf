
resource "azurerm_kubernetes_cluster" "temporal" {
  name                      = var.aks_name
  location                  = var.location
  resource_group_name       = var.rg_name
  dns_prefix                = "aksprod"
  sku_tier                  = "Free"
  kubernetes_version        = "1.28"
  automatic_channel_upgrade = "stable"
  node_resource_group       = "${var.rg_name}-${var.aks_name}-prod"
  oidc_issuer_enabled       = true
  workload_identity_enabled = true

  default_node_pool {
    name                = "default"
    node_count          = var.node_count
    min_count           = 1
    max_count           = 3
    vm_size             = "Standard_D2_v2"
    vnet_subnet_id      = var.aks_subnet_id
    enable_auto_scaling = true
    type                = "VirtualMachineScaleSets"

    node_labels = {
      role = "main"
    }
  }

  identity {
    type         = "UserAssigned"
    identity_ids = [var.user_aid_id]
  }

  network_profile {
    network_plugin = "azure"
    network_policy = "azure"
    service_cidr   = "10.0.0.0/16"
    dns_service_ip = "10.0.0.10"
  }



  tags = {
    Environment = var.env
  }

  lifecycle {
    ignore_changes = [default_node_pool[0].node_count]
  }
}




# Create a kubeconfig file that can be added to/replace ~/.kube/config
resource "local_file" "kubeconfig" {
  filename = "aks_kube_config"
  content  = azurerm_kubernetes_cluster.temporal.kube_config_raw
}


output "aks_id" {
  value = azurerm_kubernetes_cluster.temporal.id
}
