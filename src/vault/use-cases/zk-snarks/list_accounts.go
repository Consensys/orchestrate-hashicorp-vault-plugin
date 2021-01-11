package zksnarks

import (
	"context"

	apputils "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/use-cases"
	"github.com/hashicorp/vault/sdk/logical"
)

type listZksAccountsUseCase struct {
	storage logical.Storage
}

func NewListZksAccountsUseCase() usecases.ListZksAccountsUseCase {
	return &listZksAccountsUseCase{}
}

func (uc listZksAccountsUseCase) WithStorage(storage logical.Storage) usecases.ListZksAccountsUseCase {
	uc.storage = storage
	return &uc
}

// Execute gets a list of Ethereum accounts
func (uc *listZksAccountsUseCase) Execute(ctx context.Context, namespace string) ([]string, error) {
	logger := apputils.Logger(ctx).With("namespace", namespace)
	logger.Debug("listing zk-snarks accounts")

	return uc.storage.List(ctx, apputils.ComputeZksKey("", namespace))
}
