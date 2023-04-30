resource "kubernetes_deployment_v1" "api-gateway" {
  metadata {
    name      = "api-gateway-deployment"
    namespace = kubernetes_namespace_v1.wk_kanji_write.metadata.0.name

    labels = {
      app = "api-gateway-deployment"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "api-gateway-deployment"
      }
    }
    template {
      metadata {
        labels = {
          app = "api-gateway-deployment"
        }
      }
      spec {
        container {
          image             = "aeriqu/kanikaki/api-gateway:latest"
          image_pull_policy = "Never"
          name              = "api-gateway"

          port {
            container_port = 8080
          }

          env {
            name  = "JWT_SIGNING_KEY"
            value = var.jwt_signing_key
          }
        }
      }
    }
  }
}

resource "kubernetes_deployment_v1" "auth" {
  metadata {
    name      = "auth-deployment"
    namespace = kubernetes_namespace_v1.wk_kanji_write.metadata.0.name

    labels = {
      app = "auth-deployment"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "auth-deployment"
      }
    }
    template {
      metadata {
        labels = {
          app = "auth-deployment"
        }
      }
      spec {
        container {
          image             = "aeriqu/kanikaki/auth:latest"
          image_pull_policy = "Never"
          name              = "auth"

          port {
            container_port = 8080
          }

          env {
            name  = "MONGODB_USERNAME"
            value = var.mongodb_auth_username
          }

          env {
            name  = "MONGODB_PASSWORD"
            value = var.mongodb_auth_password
          }

          env {
            name  = "JWT_SIGNING_KEY"
            value = var.jwt_signing_key
          }

          env {
            name  = "CREDENTIAL_SALT"
            value = var.credential_salt
          }
        }
      }
    }
  }
}

resource "kubernetes_stateful_set_v1" "mongodb-auth" {
  metadata {
    name      = "mongodb-auth-set"
    namespace = kubernetes_namespace_v1.wk_kanji_write.metadata.0.name

    labels = {
      app = "mongodb-auth-set"
    }
  }
  spec {
    service_name = "mongodb-auth-set"
    replicas     = 1
    selector {
      match_labels = {
        app = "mongodb-auth-set"
      }
    }
    template {
      metadata {
        labels = {
          app = "mongodb-auth-set"
        }
      }
      spec {
        container {
          image             = var.mongodb_image_version
          image_pull_policy = "IfNotPresent"
          name              = "mongodb-auth"

          port {
            container_port = 27017
          }

          env {
            name  = "MONGO_INITDB_ROOT_USERNAME"
            value = var.mongodb_auth_username
          }

          env {
            name  = "MONGO_INITDB_ROOT_PASSWORD"
            value = var.mongodb_auth_password
          }

          env {
            name  = "MONGO_INITDB_DATABASE"
            value = "auth"
          }

          volume_mount {
            name       = "mongodb-auth-volume-data-db"
            mount_path = "/data/db"
          }
        }

        volume {
          name = "mongodb-auth-volume-data-db"
          persistent_volume_claim {
            claim_name = kubernetes_persistent_volume_claim_v1.mongodb-auth.metadata.0.name
          }
        }
      }
    }
  }
}

resource "kubernetes_deployment_v1" "hello-world" {
  metadata {
    name      = "hello-world-deployment"
    namespace = kubernetes_namespace_v1.wk_kanji_write.metadata.0.name

    labels = {
      app = "hello-world-deployment"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "hello-world-deployment"
      }
    }
    template {
      metadata {
        labels = {
          app = "hello-world-deployment"
        }
      }
      spec {
        container {
          image             = "aeriqu/kanikaki/hello-world:latest"
          image_pull_policy = "Never"
          name              = "hello-world"

          port {
            container_port = 8080
          }
        }
      }
    }
  }
}

resource "kubernetes_deployment_v1" "web" {
  metadata {
    name      = "web-deployment"
    namespace = kubernetes_namespace_v1.wk_kanji_write.metadata.0.name

    labels = {
      app = "web-deployment"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "web-deployment"
      }
    }
    template {
      metadata {
        labels = {
          app = "web-deployment"
        }
      }
      spec {
        container {
          image             = "aeriqu/kanikaki/web:latest"
          image_pull_policy = "Never"
          name              = "web"

          port {
            container_port = 8080
          }
        }
      }
    }
  }
}
