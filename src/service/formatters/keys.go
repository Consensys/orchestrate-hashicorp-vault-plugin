package formatters

import (
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/entities"
	"github.com/hashicorp/vault/sdk/logical"
)

func FormatKeyResponse(key *entities.Key) *logical.Response {
	return &logical.Response{
		Data: map[string]interface{}{
			"id":               key.ID,
			"curve":            key.Curve,
			"signingAlgorithm": key.Algorithm,
			"publicKey":        key.PublicKey,
			"namespace":        key.Namespace,
			"tags":             key.Tags,
		},
	}
}
