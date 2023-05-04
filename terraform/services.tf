resource "kubernetes_service_v1" "api-gateway" {
  metadata {
    name      = "api-gateway-service"
    namespace = kubernetes_namespace_v1.kanikaki.metadata.0.name
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
    namespace = kubernetes_namespace_v1.kanikaki.metadata.0.name
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

resource "kubernetes_service_v1" "kanji" {
  metadata {
    name      = "kanji-service"
    namespace = kubernetes_namespace_v1.kanikaki.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment_v1.kanji.metadata.0.labels.app
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
    namespace = kubernetes_namespace_v1.kanikaki.metadata.0.name
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

resource "kubernetes_service_v1" "mongodb-kanji" {
  metadata {
    name      = "mongodb-kanji-service"
    namespace = kubernetes_namespace_v1.kanikaki.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_stateful_set_v1.mongodb-kanji.metadata.0.labels.app
    }
    port {
      port        = 27017
      target_port = 27017
      protocol    = "TCP"
    }

    type = "ClusterIP"
  }
}

resource "kubernetes_service_v1" "wanikani" {
  metadata {
    name      = "wanikani"
    namespace = kubernetes_namespace_v1.kanikaki.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment_v1.wanikani.metadata.0.labels.app
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
    namespace = kubernetes_namespace_v1.kanikaki.metadata.0.name
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
