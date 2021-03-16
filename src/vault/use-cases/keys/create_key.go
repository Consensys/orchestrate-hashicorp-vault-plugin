package keys

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/log"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/pkg/crypto"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/entities"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/storage"
	usecases "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/use-cases"
	"github.com/consensys/quorum/common/hexutil"
	crypto2 "github.com/ethereum/go-ethereum/crypto"
	"github.com/hashicorp/vault/sdk/logical"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/errors"
)

type createKeyUseCase struct {
	storage logical.Storage
}

func NewCreateKeyUseCase() usecases.CreateKeyUseCase {
	return &createKeyUseCase{}
}

func (uc createKeyUseCase) WithStorage(storage logical.Storage) usecases.CreateKeyUseCase {
	uc.storage = storage
	return &uc
}

func (uc *createKeyUseCase) Execute(ctx context.Context, namespace, id, algo, curve, importedPrivKey string, tags map[string]string) (*entities.Key, error) {
	logger := log.FromContext(ctx).With("namespace", namespace).With("algorithm", algo).With("curve", curve)
	logger.Debug("creating new key")

	key := &entities.Key{
		Algorithm: algo,
		Curve:     curve,
		Namespace: namespace,
		Tags:      tags,
	}

	switch {
	case algo == entities.EDDSA && curve == entities.BN256:
		privKey, err := crypto.NewBN256()
		if err != nil {
			errMessage := "failed to generate key"
			logger.With("error", err).Error(errMessage)
			return nil, err
		}
		key.PrivateKey = hexutil.Encode(privKey.Bytes())
		key.PublicKey = hexutil.Encode(privKey.Public().Bytes())
	case algo == entities.ECDSA && curve == entities.Secp256k1:
		var privKey = new(ecdsa.PrivateKey)
		var err error
		if importedPrivKey == "" {
			privKey, err = crypto.NewSecp256k1()
			if err != nil {
				errMessage := "failed to generate Ethereum private key"
				logger.With("error", err).Error(errMessage)
				return nil, err
			}
		} else {
			privKey, err = crypto.ImportSecp256k1(importedPrivKey)
			if err != nil {
				errMessage := "failed to import Ethereum private key, please verify that the provided private key is valid"
				logger.With("error", err).Error(errMessage)
				return nil, err
			}
		}

		key.PrivateKey = hex.EncodeToString(crypto2.FromECDSA(privKey))
		key.PublicKey = hexutil.Encode(crypto2.FromECDSAPub(&privKey.PublicKey))
	default:
		errMessage := "invalid signing algorithm/elliptic curve combination"
		logger.Error(errMessage)
		return nil, errors.InvalidParameterError(errMessage)
	}

	err := storage.StoreJSON(ctx, uc.storage, storage.ComputeKeysStorageKey(id, key.Namespace), key)
	if err != nil {
		errMessage := "failed to store key"
		logger.With("error", err).Error(errMessage)
		return nil, err
	}

	logger.With("pub_key", key.PublicKey).Info("key pair created successfully")
	return key, nil
}
