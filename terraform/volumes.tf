resource "kubernetes_storage_class_v1" "local-storage" {
  metadata {
    name = "local-storage"
  }
  storage_provisioner = "kubernetes.io/no-provisioner"
  volume_binding_mode = "WaitForFirstConsumer"
}

resource "kubernetes_persistent_volume_claim_v1" "mongodb-auth" {
  metadata {
    name      = "mongodb-auth-volume-claim"
    namespace = kubernetes_namespace_v1.wk_kanji_write.metadata.0.name
  }
  spec {
    volume_name        = "${kubernetes_persistent_volume_v1.mongodb-auth.metadata.0.name}"
    access_modes       = ["ReadWriteOnce"]
    storage_class_name = "local-storage"
    resources {
      requests = {
        storage = kubernetes_persistent_volume_v1.mongodb-auth.spec.0.capacity.storage
      }
    }
  }
}

resource "kubernetes_persistent_volume_v1" "mongodb-auth" {
  metadata {
    name = "mongodb-auth-volume"
  }
  spec {
    access_modes                     = ["ReadWriteOnce"]
    persistent_volume_reclaim_policy = "Retain"
    volume_mode                      = "Filesystem"
    storage_class_name               = "local-storage"

    capacity = {
      storage = "2Gi"
    }

    persistent_volume_source {
      host_path {
        path = "/apps/kanikaki/volumes/mongodb-auth-volume"
        type = "DirectoryOrCreate"
      }
    }
  }
}
