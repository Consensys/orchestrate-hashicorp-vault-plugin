package ethereum

import (
	"context"
	apputils "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/use-cases"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/use-cases/ethereum/utils"
	"github.com/hashicorp/vault/sdk/logical"
)

// listAccountsUseCase is a use case to get a list of Ethereum accounts
type listAccountsUseCase struct {
	storage logical.Storage
}

// NewListAccountUseCase creates a new ListAccountsUseCase
func NewListAccountsUseCase() use_cases.ListAccountsUseCase {
	return &listAccountsUseCase{}
}

func (uc listAccountsUseCase) WithStorage(storage logical.Storage) use_cases.ListAccountsUseCase {
	uc.storage = storage
	return &uc
}

// Execute creates a new Ethereum account and stores it in the Vault
func (uc *listAccountsUseCase) Execute(ctx context.Context, namespace string) ([]string, error) {
	logger := apputils.Logger(ctx).With("namespace", namespace)
	logger.Debug("listing Ethereum accounts")

	return uc.storage.List(ctx, utils.ComputeKey("", namespace))
}
