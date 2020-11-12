package src

import (
	"context"
	builder "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/ethereum"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/service/ethereum"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

// NewVaultBackend returns the Hashicorp Vault backend
func NewVaultBackend(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	ethereumController := ethereum.NewController(builder.NewEthereumUseCases())

	vaultPlugin := &framework.Backend{
		Help:  "Orchestrate Hashicorp Vault Plugin. Please submit a request for help.",
		Paths: ethereumController.Paths(),
		PathsSpecial: &logical.Paths{
			SealWrapStorage: []string{
				"ethereum/accounts",
			},
		},
		Secrets:     []*framework.Secret{},
		BackendType: logical.TypeLogical,
	}

	ctx = utils.WithLogger(ctx, vaultPlugin.Logger())
	if err := vaultPlugin.Setup(ctx, conf); err != nil {
		return nil, err
	}

	return vaultPlugin, nil
}
