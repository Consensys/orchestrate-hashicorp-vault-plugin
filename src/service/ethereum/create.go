package ethereum

import (
	"context"
	ethereum "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/ethereum/use-cases"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/service/formatters"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

type createOperation struct {
	properties *framework.OperationProperties
	useCase    ethereum.CreateAccountUseCase
}

func NewCreateOperation(useCase ethereum.CreateAccountUseCase) framework.OperationHandler {
	exampleAccount := ExampleETHAccount()

	return &createOperation{
		properties: &framework.OperationProperties{
			Summary:     "Creates a new Ethereum account",
			Description: "Creates a new Ethereum account by generating a private key, storing it in the Vault and computing its public key and address",
			Examples: []framework.RequestExample{
				{
					Description: "Creates a new account on the tenant0 namespace",
					Data: map[string]interface{}{
						namespaceLabel: exampleAccount.Namespace,
					},
					Response: Example200Response(),
				},
			},
			Responses: map[int][]framework.Response{
				400: {Example400Response()},
				500: {Example500Response()},
			},
		},
		useCase: useCase,
	}
}

func (op *createOperation) Handler() framework.OperationFunc {
	return func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
		namespace := data.Get("namespace").(string)

		account, err := op.useCase.WithStorage(req.Storage).Execute(ctx, namespace, "")
		if err != nil {
			return nil, err
		}

		return formatters.FormatAccountResponse(account), nil
	}
}

func (op *createOperation) Properties() framework.OperationProperties {
	return *op.properties
}
