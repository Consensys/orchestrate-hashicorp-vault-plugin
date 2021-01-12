package zksnarks

import (
	"context"

	apputils "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/use-cases"
	eddsa "github.com/consensys/gnark/crypto/signature/eddsa/bn256"
	"github.com/hashicorp/vault/sdk/logical"

	"github.com/consensys/quorum/common/hexutil"
)

type signPayloadUseCase struct {
	getAccountUC usecases.GetZksAccountUseCase
}

func NewSignUseCase(getAccountUC usecases.GetZksAccountUseCase) usecases.SignUseCase {
	return &signPayloadUseCase{
		getAccountUC: getAccountUC,
	}
}

func (uc signPayloadUseCase) WithStorage(storage logical.Storage) usecases.SignUseCase {
	uc.getAccountUC = uc.getAccountUC.WithStorage(storage)
	return &uc
}

// Execute signs an arbitrary payload using an existing Ethereum account
func (uc *signPayloadUseCase) Execute(ctx context.Context, address, namespace, data string) (string, error) {
	logger := apputils.Logger(ctx).With("namespace", namespace).With("address", address)
	logger.Debug("signing message")

	account, err := uc.getAccountUC.Execute(ctx, address, namespace)
	if err != nil {
		return "", err
	}
	
	signature, err := eddsa.Sign([]byte(data), account.PublicKey, account.PrivateKey)
	if err != nil {
		errMessage := "failed to sign payload using EDDSA"
		logger.With("error", err).Error(errMessage)
		return "", err
	}

	logger.Info("payload signed successfully")
	return hexutil.Encode([]byte(signature.R.X.String())), nil
}
