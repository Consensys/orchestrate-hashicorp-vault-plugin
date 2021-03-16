package keys

import (
	"context"

	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/log"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/service/formatters"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (c *controller) NewImportOperation() *framework.PathOperation {
	exampleAccount := utils.ExampleKey()
	successExample := utils.Example200KeyResponse()

	return &framework.PathOperation{
		Callback:    c.importHandler(),
		Summary:     "Imports a key pair",
		Description: "Imports a key pair given a private key, storing it in the Vault and computing its public key and address",
		Examples: []framework.RequestExample{
			{
				Description: "Imports a key pair on the tenant0 namespace",
				Data: map[string]interface{}{
					formatters.PrivateKeyLabel: exampleAccount.PrivateKey,
				},
				Response: successExample,
			},
		},
		Responses: map[int][]framework.Response{
			200: {*successExample},
			400: {utils.Example400Response()},
			500: {utils.Example500Response()},
		},
	}
}

func (c *controller) importHandler() framework.OperationFunc {
	return func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
		namespace := formatters.GetRequestNamespace(req)
		id := data.Get(formatters.AccountIDLabel).(string)
		curve := data.Get(formatters.DataLabel).(string)
		algo := data.Get(formatters.DataLabel).(string)
		tags := data.Get(formatters.TagsLabel).(map[string]string)
		privateKeyString := data.Get(formatters.PrivateKeyLabel).(string)

		if privateKeyString == "" {
			return logical.ErrorResponse("privateKey must be provided"), nil
		}

		ctx = log.Context(ctx, c.logger)
		key, err := c.useCases.CreateKey().WithStorage(req.Storage).Execute(ctx, namespace, id, algo, curve, privateKeyString, tags)
		if err != nil {
			return nil, err
		}

		return formatters.FormatKeyResponse(key), nil
	}
}
