terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "2.37.1"
    }
  }
}

provider "kubernetes" {
  config_path    = var.kubernetes_config_path
  config_context = var.kubernetes_config_context
}

resource "kubernetes_namespace_v1" "kanikaki" {
  metadata {
    name = "kanikaki"
  }
}