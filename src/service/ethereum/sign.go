package ethereum

import (
	"context"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/service/errors"

	errors2 "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/pkg/errors"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/pkg/log"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/service/formatters"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (c *controller) NewSignPayloadOperation() *framework.PathOperation {
	exampleAccount := utils.ExampleETHAccount()

	return &framework.PathOperation{
		Callback:    c.signPayloadHandler(),
		Summary:     "Signs an arbitrary message using an existing Ethereum account",
		Description: "Signs an arbitrary message using ECDSA and the private key of an existing Ethereum account",
		Examples: []framework.RequestExample{
			{
				Description: "Signs a message",
				Data: map[string]interface{}{
					formatters.IDLabel:   exampleAccount.Address,
					formatters.DataLabel: "my data to sign",
				},
				Response: utils.Example200ResponseSignature(),
			},
		},
		Responses: map[int][]framework.Response{
			200: {*utils.Example200ResponseSignature()},
			400: {utils.Example400Response()},
			404: {utils.Example404Response()},
			500: {utils.Example500Response()},
		},
	}
}

func (c *controller) signPayloadHandler() framework.OperationFunc {
	return func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
		address := data.Get(formatters.IDLabel).(string)
		payload := data.Get(formatters.DataLabel).(string)
		namespace := formatters.GetRequestNamespace(req)

		if payload == "" {
			return errors.ParseHTTPError(errors2.InvalidFormatError("payload must be provided"))
		}

		ctx = log.Context(ctx, c.logger)
		signature, err := c.useCases.SignPayload().WithStorage(req.Storage).Execute(ctx, address, namespace, payload)
		if err != nil {
			return errors.ParseHTTPError(err)
		}

		return formatters.FormatSignatureResponse(signature), nil
	}
}
