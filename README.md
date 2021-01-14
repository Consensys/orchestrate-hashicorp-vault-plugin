[![Website](https://img.shields.io/website?label=documentation&url=https%3A%2F%2Fdocs.orchestrate.pegasys.tech%2F)](https://docs.orchestrate.pegasys.tech/)
[![Website](https://img.shields.io/website?label=website&url=https%3A%2F%2Fpegasys.tech%2Forchestrate%2F)](https://pegasys.tech/orchestrate/)
[![CircleCI](https://img.shields.io/circleci/build/gh/ConsenSys/orchestrate-hashicorp-vault-plugin?token=8a52ab8f0640f5bee56991cd30d808f735749dbf)](https://circleci.com/gh/PegaSysEng/orchestrate-hashicorp-vault-plugin)

![](/img/codefi_orchestrate_logo.png)

# Orchestrate Hashicorp Vault plugin

The Orchestrate library provides a secure Hashicorp Vault plugin to store your keys and perform
 cryptographic operations such as:
 - Ethereum transactions
 - Besu-Orion private transaction
 - Quorum-Tessera private transaction
 - ZKP signing operation

## Compatibility

| SDK versions        | Orchestrate versions       |
| ------------------- | -------------------------- |
| master/HEAD         | v21.1.0 (Unreleased)		   |
| Plugin v0.6.0       | v21.1.0 (Unreleased) 		   |

## Running in development

### Pre-requirements
- Go >= 1.14
- Makefile
- docker-compose

First we start up Hashicorp server with our Orchestrate plugin in dev mode
```
$ make dev
```

If everything went well you should see the following output:
```
...
vault-dev_1       | WARNING! dev mode is enabled! In this mode, Vault runs entirely in-memory
vault-dev_1       | and starts unsealed with a single unseal key. The root token is already
vault-dev_1       | authenticated to the CLI, so you can immediately begin using Vault.
vault-dev_1       | 
vault-dev_1       | You may need to set the following environment variable:
vault-dev_1       | 
vault-dev_1       |     $ export VAULT_ADDR='http://0.0.0.0:8200'
vault-dev_1       | 
vault-dev_1       | The unseal key and root token are displayed below in case you want to
vault-dev_1       | seal/unseal the Vault or re-authenticate.
vault-dev_1       | 
vault-dev_1       | Unseal Key: 3hTanFX/q99PMBubXjwL/cFXh3YKCABPSmw31Jwok1w=
vault-dev_1       | Root Token: DevVaultToken
vault-dev_1       | 
vault-dev_1       | The following dev plugins are registered in the catalog:
vault-dev_1       |     - orchestrate
```

To stop our environment we should run:
```
$ make down
```

## Contributing
[How to Contribute](CONTRIBUTING.md)

## Orchestrate documentation

For a global understanding of Orchestrate, not only this Hashicorp Vault plugin, refer to the
[Orchestrate documentation.](https://docs.orchestrate.consensys.net/)
