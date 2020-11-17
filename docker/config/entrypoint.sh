#!/bin/bash

CONFIG_DIR="/home/vault/config"
INIT_SCRIPT="/home/vault/config/init.sh"

export VAULT_ADDR="http://localhost:9200"

mkdir -p $CONFIG_DIR

nohup vault server -config /home/vault/config/vault.hcl &
VAULT_PID=$!

if [ -f "$INIT_SCRIPT" ]; then
  /bin/bash $INIT_SCRIPT
fi

wait $VAULT_PID
