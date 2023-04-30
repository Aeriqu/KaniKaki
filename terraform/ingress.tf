resource "kubernetes_ingress_v1" "nginx" {
  wait_for_load_balancer = true
  metadata {
    name      = "nginx-ingress"
    namespace = kubernetes_namespace_v1.wk_kanji_write.metadata.0.name
    // annotations are intended to allow cors for the local development environment
    // so complete docker rebuilds and redeploys aren't required. since it's localhost
    // this shouldn't pose too much of a risk.
    annotations = {
      "nginx.ingress.kubernetes.io/enable-cors" = "true"
      "nginx.ingress.kubernetes.io/cors-allow-origin" = "http://localhost:3000"
      "nginx.ingress.kubernetes.io/cors-allow-methods" = "PUT, GET, POST, OPTIONS, DELETE"
    }
  }
  spec {
    ingress_class_name = "nginx"
    rule {
      http {
        path {
          path = "/api"
          backend {
            service {
              name = kubernetes_service_v1.api-gateway.metadata.0.name
              port {
                number = 80
              }
            }
          }
        }

        path {
          path = "/"
          backend {
            service {
              name = kubernetes_service_v1.web.metadata.0.name
              port {
                number = 80
              }
            }
          }
        }
      }
    }
  }
}