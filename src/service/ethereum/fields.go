package ethereum

import "github.com/hashicorp/vault/sdk/framework"

const (
	privateKeyLabel = "privateKey"
	addressLabel    = "address"

	namespaceHeader = "X-Vault-Namespace"
)

var namespaceFieldSchema = &framework.FieldSchema{
	Type:        framework.TypeString,
	Description: "Namespace in which to store the account",
	Required:    false,
	Default:     "",
}

var addressFieldSchema = &framework.FieldSchema{
	Type:        framework.TypeString,
	Description: "Address of the account",
	Required:    true,
}
