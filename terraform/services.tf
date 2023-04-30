resource "kubernetes_service_v1" "api-gateway" {
  metadata {
    name      = "api-gateway-service"
    namespace = kubernetes_namespace_v1.wk_kanji_write.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment_v1.api-gateway.metadata.0.labels.app
    }
    port {
      port        = 80
      target_port = 8080
      protocol    = "TCP"
    }

    type = "ClusterIP"
  }
}

resource "kubernetes_service_v1" "auth" {
  metadata {
    name      = "auth-service"
    namespace = kubernetes_namespace_v1.wk_kanji_write.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment_v1.auth.metadata.0.labels.app
    }
    port {
      port        = 80
      target_port = 8080
      protocol    = "TCP"
    }

    type = "ClusterIP"
  }
}

resource "kubernetes_service_v1" "mongodb-auth" {
  metadata {
    name      = "mongodb-auth-service"
    namespace = kubernetes_namespace_v1.wk_kanji_write.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_stateful_set_v1.mongodb-auth.metadata.0.labels.app
    }
    port {
      port        = 27017
      target_port = 27017
      protocol    = "TCP"
    }

    type = "ClusterIP"
  }
}

resource "kubernetes_service_v1" "hello-world" {
  metadata {
    name      = "hello-world-service"
    namespace = kubernetes_namespace_v1.wk_kanji_write.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment_v1.hello-world.metadata.0.labels.app
    }
    port {
      port        = 80
      target_port = 8080
      protocol    = "TCP"
    }

    type = "ClusterIP"
  }
}

resource "kubernetes_service_v1" "web" {
  metadata {
    name      = "web"
    namespace = kubernetes_namespace_v1.wk_kanji_write.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment_v1.web.metadata.0.labels.app
    }
    port {
      port        = 80
      target_port = 8080
      protocol    = "TCP"
    }

    type = "ClusterIP"
  }
}
