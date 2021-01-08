package usecases

import (
	"context"

	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/entities"
	"github.com/hashicorp/vault/sdk/logical"
)

//go:generate mockgen -source=zk-snarks.go -destination=mocks/zk-snarks.go -package=mocks


type ZkSnarksUseCases interface {
	CreateAccount() CreateAccountUseCase
}

type CreateBN256AccountUseCase interface {
	Execute(ctx context.Context, namespace string) (*entities.ZkSnarksAccount, error)
	WithStorage(storage logical.Storage) CreateBN256AccountUseCase
}
