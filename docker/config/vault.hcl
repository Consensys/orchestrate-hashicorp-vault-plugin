default_lease_ttl = "24h"
disable_mlock = "true"
max_lease_ttl = "43800h"
api_addr = "http://localhost:9200"
ui = "true"
plugin_directory = "/home/vault/plugins"
log_level = "Debug"

backend "file" {
  path = "/home/vault/config/data"
}

listener "tcp" {
  address = "0.0.0.0:9200"
  tls_disable = true
}