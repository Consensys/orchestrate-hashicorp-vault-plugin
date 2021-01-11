package builder

import (
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/use-cases"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/use-cases/zk-snarks"
)

type zkSnarksUseCases struct {
	createAccount usecases.CreateZksAccountUseCase
}

func NewZkSnarksUseCases() usecases.ZksUseCases {
	return &zkSnarksUseCases{
		createAccount: zksnarks.NewCreateAccountUseCase(),
	}
}

func (z *zkSnarksUseCases) CreateAccount() usecases.CreateZksAccountUseCase {
	return z.createAccount
}
