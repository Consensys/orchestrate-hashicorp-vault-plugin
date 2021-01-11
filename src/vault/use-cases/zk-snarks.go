package usecases

import (
	"context"

	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/entities"
	"github.com/hashicorp/vault/sdk/logical"
)

//go:generate mockgen -source=zk-snarks.go -destination=mocks/zk-snarks.go -package=mocks


type ZksUseCases interface {
	CreateAccount() CreateZksAccountUseCase
}

type CreateZksAccountUseCase interface {
	Execute(ctx context.Context, namespace string) (*entities.ZksAccount, error)
	WithStorage(storage logical.Storage) CreateZksAccountUseCase
}
