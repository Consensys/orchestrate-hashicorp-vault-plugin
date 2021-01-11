package zksnarks

import (
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
			// logical.ListOperation:   c.NewListOperation(),
			// logical.ReadOperation:   c.NewListOperation(),
		},
	}
}

func getNamespace(req *logical.Request) string {
	namespace := ""

	if val, hasVal := req.Headers[formatters.NamespaceHeader]; hasVal {
		namespace = val[0]
	}

	return namespace
}
