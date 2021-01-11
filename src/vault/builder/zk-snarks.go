package builder

import (
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/use-cases"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/use-cases/zk-snarks"
)

type zkSnarksUseCases struct {
	createAccount usecases.CreateZksAccountUseCase
	getAccount    usecases.GetZksAccountUseCase
	listAccounts  usecases.ListZksAccountsUseCase
}

func NewZkSnarksUseCases() usecases.ZksUseCases {
	return &zkSnarksUseCases{
		createAccount: zksnarks.NewCreateAccountUseCase(),
		getAccount:    zksnarks.NewGetZksAccountUseCase(),
		listAccounts:  zksnarks.NewListZksAccountsUseCase(),
	}
}

func (z *zkSnarksUseCases) CreateAccount() usecases.CreateZksAccountUseCase {
	return z.createAccount
}

func (z *zkSnarksUseCases) GetAccount() usecases.GetZksAccountUseCase {
	return z.getAccount
}

func (z *zkSnarksUseCases) ListAccounts() usecases.ListZksAccountsUseCase {
	return z.listAccounts
}
