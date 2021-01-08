package ethereum

import (
	"context"

	apputils "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/entities"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/storage"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/use-cases"
	"github.com/consensys/gnark/crypto/hash/mimc/bn256"
	eddsa "github.com/consensys/gnark/crypto/signature/eddsa/bn256"
	"github.com/hashicorp/vault/sdk/logical"
)

type createBN256AccountUseCase struct {
	storage logical.Storage
}

func NewCreateBN256AccountUseCase() usecases.CreateBN256AccountUseCase {
	return &createBN256AccountUseCase{}
}

func (uc createBN256AccountUseCase) WithStorage(storage logical.Storage) usecases.CreateBN256AccountUseCase {
	uc.storage = storage
	return &uc
}

func (uc *createBN256AccountUseCase) Execute(ctx context.Context, namespace string) (*entities.ZkSnarksAccount, error) {
	logger := apputils.Logger(ctx).With("namespace", namespace)
	logger.Debug("creating new zk-snarks bn256 account")

	// @TODO Generate random seed
	var seed [32]byte
	s := []byte("eddsa")
	for i, v := range s {
		seed[i] = v
	}

	hFunc := bn256.NewMiMC("seed")
	pubKey, privKey := eddsa.New(seed, hFunc)
	account := &entities.ZkSnarksAccount{
		PrivateKey: privKey,
		Address:    pubKey.A.X.String(),
		PublicKey:  pubKey,
		Namespace:  namespace,
	}

	err := storage.StoreJSON(ctx, uc.storage, apputils.ComputeZkSnarkKey(account.Address, account.Namespace), account)
	if err != nil {
		errMessage := "failed to create account entry"
		apputils.Logger(ctx).With("error", err).Error(errMessage)
		return nil, err
	}

	logger.With("address", account.Address).Info("zk-snarks bn256 account created successfully")
	return account, nil
}
