package ethereum

import (
	"context"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/service/formatters"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (c *controller) NewSignPayloadOperation() *framework.PathOperation {
	exampleAccount := utils.ExampleETHAccount()
	successExample := utils.Example200Response()

	return &framework.PathOperation{
		Callback:    c.signPayloadHandler(),
		Summary:     "Signs an arbitrary message using an existing Ethereum account",
		Description: "Signs an arbitrary message using ECDSA and the private key of an existing Ethereum account",
		Examples: []framework.RequestExample{
			{
				Description: "Signs a message",
				Data: map[string]interface{}{
					addressLabel: exampleAccount.Address,
					dataLabel:    "my data to sign",
				},
				Response: successExample,
			},
		},
		Responses: map[int][]framework.Response{
			200: {framework.Response{
				Description: "Success",
				Example: &logical.Response{
					Data: map[string]interface{}{
						"signature": "0x7107193a8683e258ada2dfa76b5e6fc145ebd98f0e6eee77cb91381201fe7bca5445beccebe164e23abe0639f089e17b24ce867be9fece8b4872cfe13d91464601",
					},
				},
			}},
			400: {utils.Example400Response()},
			500: {utils.Example500Response()},
		},
	}
}

func (c *controller) signPayloadHandler() framework.OperationFunc {
	return func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
		address := data.Get("address").(string)
		payload := data.Get("data").(string)
		namespace := getNamespace(req)

		if payload == "" {
			return logical.ErrorResponse("data must be provided"), nil
		}

		ctx = utils.WithLogger(ctx, c.logger)
		signature, err := c.useCases.SignPayload().WithStorage(req.Storage).Execute(ctx, address, namespace, payload)
		if err != nil {
			return nil, err
		}

		return formatters.FormatSignatureResponse(signature), nil
	}
}
