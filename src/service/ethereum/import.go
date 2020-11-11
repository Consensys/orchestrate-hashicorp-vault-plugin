package ethereum

import (
	"context"
	ethereum "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/ethereum/use-cases"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

type importOperation struct {
	properties *framework.OperationProperties
	useCase    ethereum.CreateAccountUseCase
}

func NewImportOperation(useCase ethereum.CreateAccountUseCase) framework.OperationHandler {
	exampleAccount := ExampleETHAccount()

	return &importOperation{
		properties: &framework.OperationProperties{
			Summary:     "Imports an Ethereum account",
			Description: "Imports an Ethereum account given a private key, storing it in the Vault and computing its public key and address",
			Examples: []framework.RequestExample{
				{
					Description: "Creates a new account on the tenant0 namespace",
					Data: map[string]interface{}{
						namespaceLabel:  exampleAccount.Namespace,
						privateKeyLabel: exampleAccount.PrivateKey,
					},
					Response: Example200Response(),
				},
			},
			Responses: map[int][]framework.Response{
				400: {Example400Response()},
				422: {Example422Response()},
				500: {Example500Response()},
			},
		},
		useCase: useCase,
	}
}

func (op *importOperation) Handler() framework.OperationFunc {
	return func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
		privateKeyString := data.Get("privateKey").(string)
		namespace := data.Get("namespace").(string)

		account, err := op.useCase.WithStorage(req.Storage).Execute(ctx, namespace, privateKeyString)
		if err != nil {
			return nil, err
		}

		return FormatAccountResponse(account), nil
	}
}

func (op *importOperation) Properties() framework.OperationProperties {
	return *op.properties
}
