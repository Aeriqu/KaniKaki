variable "kubernetes_config_path" {
  type = string
}

variable "kubernetes_config_context" {
  type = string
}

variable "mongodb_image_version" {
  type = string
}

variable "mongodb_auth_username" {
  type = string
  sensitive = true
}

variable "mongodb_auth_password" {
  type = string
  sensitive = true
}

variable "mongodb_kanji_username" {
  type = string
  sensitive = true
}

variable "mongodb_kanji_password" {
  type = string
  sensitive = true
}

variable "credential_salt" {
  type = string
  sensitive = true
}

variable "jwt_signing_key" {
  type = string
  sensitive = true
}