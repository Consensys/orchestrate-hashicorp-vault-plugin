package main

import (
	"github.com/hashicorp/go-hclog"
	"os"

	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/sdk/plugin"
)

func main() {
	client := &api.PluginAPIClientMeta{}
	logger := hclog.New(&hclog.LoggerOptions{})
	err := client.FlagSet().Parse(os.Args[1:])
	if err != nil {
		logger.Error("failed to parse args", "error", err)
		os.Exit(1)
	}

	err = plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: src.NewVaultBackend,
		TLSProviderFunc:    api.VaultPluginTLSProvider(client.GetTLSConfig()),
		Logger:             logger,
	})
	if err != nil {
		logger.Error("failed to start plugin", "error", err)
		os.Exit(1)
	}
}
