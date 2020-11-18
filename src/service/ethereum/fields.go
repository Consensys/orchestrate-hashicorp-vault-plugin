package ethereum

import "github.com/hashicorp/vault/sdk/framework"

const (
	namespaceLabel  = "namespace"
	privateKeyLabel = "privateKey"
	addressLabel    = "address"
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
