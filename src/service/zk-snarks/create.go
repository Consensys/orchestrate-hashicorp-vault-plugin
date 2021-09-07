package zksnarks

import (
	"context"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/service/errors"

	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/pkg/log"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/service/formatters"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (c *controller) NewCreateOperation() *framework.PathOperation {
	successExample := utils.Example200ZksResponse()

	return &framework.PathOperation{
		Callback:    c.createHandler(),
		Summary:     "Creates a new zk-snarks account",
		Description: "Creates a new zk-snarks account by generating a private key, storing it in the Vault and computing its public key and address",
		Examples: []framework.RequestExample{
			{
				Description: "Creates a new account on the tenant0 namespace",
				Response:    successExample,
			},
		},
		Responses: map[int][]framework.Response{
			200: {*successExample},
			500: {utils.Example500Response()},
		},
	}
}

func (c *controller) createHandler() framework.OperationFunc {
	return func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
		namespace := formatters.GetRequestNamespace(req)

		ctx = log.Context(ctx, c.logger)
		account, err := c.useCases.CreateAccount().WithStorage(req.Storage).Execute(ctx, namespace)
		if err != nil {
			return errors.ParseHTTPError(err)
		}

		return formatters.FormatZksAccountResponse(account), nil
	}
}
