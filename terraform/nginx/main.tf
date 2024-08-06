resource "helm_release" "ingress_nginx" {
  name             = "nginx-ingress"
  repository       = "https://kubernetes.github.io/ingress-nginx"
  chart            = "ingress-nginx"
  namespace        = "ingress"
  create_namespace = true

  set {
    name  = "controller.service.loadBalancerIP"
    value = var.temporal_pip
  }
  set {
    name  = "controller.service.externalTrafficPolicy"
    value = "Local"
  }

  lifecycle {
    ignore_changes = [
      set,
    ]
  }
}
