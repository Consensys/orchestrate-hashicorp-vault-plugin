package zksnarks

import (
	"context"

	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/log"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/entities"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/storage"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/use-cases"
	"github.com/consensys/gnark/crypto/hash/mimc/bn256"
	eddsa "github.com/consensys/gnark/crypto/signature/eddsa/bn256"
	"github.com/hashicorp/vault/sdk/logical"
)

type createAccountUseCase struct {
	storage logical.Storage
}

func NewCreateAccountUseCase() usecases.CreateZksAccountUseCase {
	return &createAccountUseCase{}
}

func (uc createAccountUseCase) WithStorage(storage logical.Storage) usecases.CreateZksAccountUseCase {
	uc.storage = storage
	return &uc
}

func (uc *createAccountUseCase) Execute(ctx context.Context, namespace string) (*entities.ZksAccount, error) {
	logger := log.FromContext(ctx).With("namespace", namespace)
	logger.Debug("creating new zk-snarks bn256 account")

	// @TODO Generate random seed
	var seed [32]byte
	s := []byte("eddsa")
	for i, v := range s {
		seed[i] = v
	}

	hFunc := bn256.NewMiMC("seed")
	pubKey, _ := eddsa.New(seed, hFunc)
	account := &entities.ZksAccount{
		Algorithm:  entities.ZksAlgorithmEDDSA,
		Curve:      entities.ZksCurveBN256,
		PrivateKey: "TODO", // @TODO
		Address:    pubKey.A.X.String(), // @TODO
		PublicKey:  pubKey.A.X.String(), // @TODO
		Namespace:  namespace,
	}

	err := storage.StoreJSON(ctx, uc.storage, storage.ComputeZksStorageKey(account.Address, account.Namespace), account)
	if err != nil {
		errMessage := "failed to create account entry"
		logger.With("error", err).Error(errMessage)
		return nil, err
	}

	logger.With("address", account.Address).Info("zk-snarks bn256 account created successfully")
	return account, nil
}
