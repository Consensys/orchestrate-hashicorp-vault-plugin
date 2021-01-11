package formatters

import (
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/entities"
	"github.com/hashicorp/vault/sdk/logical"
)

func FormatZksAccountResponse(account *entities.ZksAccount) *logical.Response {
	return &logical.Response{
		Data: map[string]interface{}{
			"curve":      account.Curve,
			"algorithm":  account.Algorithm,
			"address":    account.Address,
			"publicKey":  account.PublicKey,
			"privateKey": account.PrivateKey,
			"namespace":  account.Namespace,
		},
	}
}
