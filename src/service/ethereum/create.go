package ethereum

import (
	"context"
	ethereum "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/ethereum/use-cases"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

type createOperation struct {
	properties *framework.OperationProperties
	useCase    ethereum.CreateAccountUseCase
}

func NewCreateOperation(useCase ethereum.CreateAccountUseCase) framework.OperationHandler {
	return &createOperation{
		properties: &framework.OperationProperties{
			Summary:     "Imports an Ethereum account given a private key",
			Description: "",
			Examples: []framework.RequestExample{
				{
					Description: "",
					Data:        nil,
					Response:    nil,
				},
			},
			Responses: map[int][]framework.Response{
				400: nil,
				422: nil,
				500: nil,
			},
		},
		useCase: useCase,
	}
}

func (op *createOperation) Handler() framework.OperationFunc {
	return func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
		namespace := data.Get("namespace").(string)

		account, err := op.useCase.Execute(ctx, namespace, "")
		if err != nil {
			// b.Logger().Error("Failed to save the new account to storage", "error", err)
			return nil, err
		}

		return FormatAccountResponse(account), nil
	}
}

func (op *createOperation) Properties() framework.OperationProperties {
	return *op.properties
}
