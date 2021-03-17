package keys

import (
	"context"

	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/log"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/service/errors"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/service/formatters"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (c *controller) NewCreateOperation() *framework.PathOperation {
	successExample := utils.Example200KeysResponse()

	return &framework.PathOperation{
		Callback:    c.createHandler(),
		Summary:     "Creates a new key pair",
		Description: "Creates a new key pair by generating a private key, storing it in the Vault and computing its public key",
		Examples: []framework.RequestExample{
			{
				Description: "Creates a new key pair on the tenant0 namespace",
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
		id := data.Get(formatters.IDLabel).(string)
		curve := data.Get(formatters.CurveLabel).(string)
		algo := data.Get(formatters.AlgoLabel).(string)
		tags := data.Get(formatters.TagsLabel).(map[string]string)

		if id == "" {
			return errors.BadRequestError(req, "id must be provided")
		}
		if curve == "" {
			return errors.BadRequestError(req, "curve must be provided")
		}
		if algo == "" {
			return errors.BadRequestError(req, "algorithm must be provided")
		}

		ctx = log.Context(ctx, c.logger)
		key, err := c.useCases.CreateKey().WithStorage(req.Storage).Execute(ctx, namespace, id, algo, curve, "", tags)
		if err != nil {
			return nil, err
		}

		return formatters.FormatKeyResponse(key), nil
	}
}