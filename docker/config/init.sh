#!/bin/bash

OPERATOR_JSON="/home/vault/config/operator.json"

function banner() {
  echo "+----------------------------------------------------------------------------------+"
  printf "| %-80s |\n" "$(date)"
  echo "|                                                                                  |"
  printf "| %-80s |\n" "$@"
  echo "+----------------------------------------------------------------------------------+"
}

function authenticate() {
  banner "Authenticating to $VAULT_ADDR as root"
  ROOT=$(echo $OPERATOR_SECRETS | jq -r .root_token)
  export VAULT_TOKEN=$ROOT
}

function unauthenticate() {
  banner "Unsetting VAULT_TOKEN"
  unset VAULT_TOKEN
}

function unseal() {
  banner "Unsealing $VAULT_ADDR..."
  UNSEAL=$(echo $OPERATOR_SECRETS | jq -r '.unseal_keys_hex[0]')
  vault operator unseal $UNSEAL
}

function configure() {
  banner "Installing orchestrate-hashicorp-vault plugin at $VAULT_ADDR..."
  SHA256PLUGIN=$(cat /home/vault/plugins/SHA256PLUGIN | awk '{print $1}')
  vault write sys/plugins/catalog/secret/orchestrate \
    sha_256="$SHA256PLUGIN" \
    command="orchestrate-hashicorp-vault-plugin"

  if [[ $? -eq 2 ]]; then
    echo "orchestrate-hashicorp-vault-plugin couldn't be written to the catalog!"
    exit 2
  fi

  vault secrets enable -path=orchestrate -plugin-name=orchestrate plugin
  if [[ $? -eq 2 ]]; then
    echo "orchestrate-hashicorp-vault-plugin couldn't be enabled!"
    exit 2
  fi
  vault audit enable file file_path=stdout
}

function status() {
  vault status
}

function init() {
  OPERATOR_SECRETS=$(vault operator init -key-shares=1 -key-threshold=1 -format=json | jq .)
  echo $OPERATOR_SECRETS >$OPERATOR_JSON
}

init
unseal
authenticate
configure
unauthenticate
