package zksnarks

import (
	"context"
	apputils "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/entities"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/storage"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/use-cases"
	"github.com/hashicorp/vault/sdk/logical"
)

type getZksAccountUseCase struct {
	storage logical.Storage
}

func NewGetZksAccountUseCase() usecases.GetZksAccountUseCase {
	return &getZksAccountUseCase{}
}

func (uc getZksAccountUseCase) WithStorage(storage logical.Storage) usecases.GetZksAccountUseCase {
	uc.storage = storage
	return &uc
}

func (uc *getZksAccountUseCase) Execute(ctx context.Context, address, namespace string) (*entities.ZksAccount, error) {
	logger := apputils.Logger(ctx).With("namespace", namespace).With("address", address)
	logger.Debug("getting zk-snarks account")

	account := &entities.ZksAccount{}
	err := storage.GetJSON(ctx, uc.storage, storage.ComputeZkSnarksStorageKey(address, namespace), account)
	if err != nil {
		apputils.Logger(ctx).With("error", err).Error("failed to retrieve account from vault")
		return nil, err
	}

	logger.Debug("zk-snarks account found successfully")
	return account, nil
}
