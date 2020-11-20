package formatters

import (
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"math/big"
)

func FormatAccountResponse(account *entities.ETHAccount) *logical.Response {
	return &logical.Response{
		Data: map[string]interface{}{
			"address":             account.Address,
			"publicKey":           account.PublicKey,
			"compressedPublicKey": account.CompressedPublicKey,
			"namespace":           account.Namespace,
		},
	}
}

func FormatSignatureResponse(signature string) *logical.Response {
	return &logical.Response{
		Data: map[string]interface{}{
			"signature": signature,
		},
	}
}

func FormatSignETHTransactionRequest(requestData *framework.FieldData) (*types.Transaction, error) {
	amount, ok := new(big.Int).SetString(requestData.Get(AmountLabel).(string), 10)
	if !ok {
		return nil, logical.CodedError(400, "invalid amount")
	}

	gasPrice, ok := new(big.Int).SetString(requestData.Get(GasPriceLabel).(string), 10)
	if !ok {
		return nil, logical.CodedError(400, "invalid gas price")
	}

	data, err := hexutil.Decode(requestData.Get(DataLabel).(string))
	if err != nil {
		return nil, logical.CodedError(400, "invalid data")
	}

	nonce := requestData.Get(NonceLabel).(int)
	gasLimit := requestData.Get(GasLimitLabel).(int)
	to := requestData.Get(ToLabel).(string)
	if to == "" {
		return types.NewContractCreation(uint64(nonce), amount, uint64(gasLimit), gasPrice, data), nil
	}

	return types.NewTransaction(uint64(nonce), common.HexToAddress(to), amount, uint64(gasLimit), gasPrice, data), nil
}
