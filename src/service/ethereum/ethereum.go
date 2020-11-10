package ethereum

import (
	ethereum "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/ethereum/use-cases"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

const (
	namespaceLabel  = "namespace"
	privateKeyLabel = "privateKey"
)

var namespaceFieldSchema = &framework.FieldSchema{
	Type:        framework.TypeString,
	Description: "Namespace in which to store the account",
	Required:    false,
	Default:     "",
}

type ethereumController struct {
	useCases ethereum.UseCases
}

func NewEthereumController(useCases ethereum.UseCases) *ethereumController {
	return &ethereumController{
		useCases: useCases,
	}
}

// Paths returns the list of paths
func (c *ethereumController) Paths() []*framework.Path {
	return framework.PathAppend(
		[]*framework.Path{
			c.pathAccounts(),
			c.pathImportAccount(),
		},
	)
}

func (c *ethereumController) pathAccounts() *framework.Path {
	return &framework.Path{
		Pattern:      "ethereum/accounts",
		HelpSynopsis: "Creates a new Ethereum account",
		Fields: map[string]*framework.FieldSchema{
			namespaceLabel: namespaceFieldSchema,
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: NewCreateOperation(c.useCases.CreateAccount()),
		},
	}
}

func (c *ethereumController) pathImportAccount() *framework.Path {
	return &framework.Path{
		Pattern: "ethereum/accounts/import",
		Fields: map[string]*framework.FieldSchema{
			privateKeyLabel: {
				Type:        framework.TypeString,
				Description: "Private key in hexadecimal format",
				Required:    true,
			},
			namespaceLabel: namespaceFieldSchema,
		},
		HelpSynopsis: "Imports an Ethereum account",
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: NewImportOperation(c.useCases.CreateAccount()),
		},
	}
}
