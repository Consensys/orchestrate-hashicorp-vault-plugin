package zksnarks

import (
	"fmt"

	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/service/formatters"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/use-cases"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

type controller struct {
	useCases usecases.ZksUseCases
	logger   hclog.Logger
}

func NewController(useCases usecases.ZksUseCases, logger hclog.Logger) *controller {
	return &controller{
		useCases: useCases,
		logger:   logger,
	}
}

// Paths returns the list of paths
func (c *controller) Paths() []*framework.Path {
	return framework.PathAppend(
		[]*framework.Path{
			c.pathAccounts(),
			c.pathAccount(),
			c.pathNamespaces(),
			c.pathSignPayload(),
		},
	)
}

func (c *controller) pathAccounts() *framework.Path {
	return &framework.Path{
		Pattern:      "zk-snarks/accounts/?",
		HelpSynopsis: "Creates a new zk-snarks account or list them",
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: c.NewCreateOperation(),
			logical.UpdateOperation: c.NewCreateOperation(),
			logical.ListOperation:   c.NewListOperation(),
			logical.ReadOperation:   c.NewListOperation(),
		},
	}
}

func (c *controller) pathAccount() *framework.Path {
	return &framework.Path{
		Pattern:      fmt.Sprintf("zk-snarks/accounts/%s", framework.GenericNameRegex("address")),
		HelpSynopsis: "Get, update or delete an Ethereum account",
		Fields: map[string]*framework.FieldSchema{
			formatters.AddressLabel: formatters.AddressFieldSchema,
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ReadOperation: c.NewGetOperation(),
		},
	}
}

func (c *controller) pathNamespaces() *framework.Path {
	return &framework.Path{
		Pattern:      "ethereum/namespaces/?",
		HelpSynopsis: "Lists all ethereum namespaces",
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ListOperation: c.NewListNamespacesOperation(),
			logical.ReadOperation: c.NewListNamespacesOperation(),
		},
	}
}

func (c *controller) pathSignPayload() *framework.Path {
	return &framework.Path{
		Pattern: fmt.Sprintf("ethereum/accounts/%s/sign", framework.GenericNameRegex("address")),
		Fields: map[string]*framework.FieldSchema{
			formatters.AddressLabel: formatters.AddressFieldSchema,
			formatters.DataLabel: {
				Type:        framework.TypeString,
				Description: "data to sign",
				Required:    true,
			},
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: c.NewSignPayloadOperation(),
			logical.UpdateOperation: c.NewSignPayloadOperation(),
		},
		HelpSynopsis: "Signs an arbitrary message using an existing Ethereum account",
	}
}
