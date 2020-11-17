package ethereum

import (
	"context"
	"fmt"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/ethereum/testutils"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/ethereum/use-cases/mocks"
	mocks2 "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/service/testutils/mocks"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateOperation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	createAccountUC := mocks.NewMockCreateAccountUseCase(ctrl)
	operation := NewCreateOperation(createAccountUC)
	storage := mocks2.NewMockStorage(ctrl)
	ctx := context.Background()

	t.Run("should define correct properties", func(t *testing.T) {
		properties := operation.Properties()

		assert.NotEmpty(t, properties.Description)
		assert.NotEmpty(t, properties.Summary)
		assert.NotEmpty(t, properties.Examples[0].Description)
		assert.NotEmpty(t, properties.Examples[0].Data)
		assert.NotEmpty(t, properties.Examples[0].Response)
		assert.NotEmpty(t, properties.Responses[200])
		assert.NotEmpty(t, properties.Responses[400])
		assert.NotEmpty(t, properties.Responses[500])
	})

	t.Run("handler should execute the correct use case", func(t *testing.T) {
		account := testutils.FakeETHAccount()
		request := &logical.Request{
			Storage: storage,
		}
		data := &framework.FieldData{
			Raw: map[string]interface{}{
				namespaceLabel: account.Namespace,
			},
			Schema: map[string]*framework.FieldSchema{
				namespaceLabel: namespaceFieldSchema,
			},
		}

		createAccountUC.EXPECT().WithStorage(storage).Return(createAccountUC)
		createAccountUC.EXPECT().Execute(ctx, account.Namespace, "").Return(account, nil)

		response, err := operation.Handler()(ctx, request, data)

		assert.NoError(t, err)
		assert.Equal(t, account.Address, response.Data["address"])
		assert.Equal(t, account.PublicKey, response.Data["publicKey"])
		assert.Equal(t, account.CompressedPublicKey, response.Data["compressedPublicKey"])
		assert.Equal(t, account.Namespace, response.Data["namespace"])
	})

	t.Run("should return same error if use case fails", func(t *testing.T) {
		request := &logical.Request{
			Storage: storage,
		}
		data := &framework.FieldData{
			Raw: map[string]interface{}{
				namespaceLabel: "myNamespace",
			},
			Schema: map[string]*framework.FieldSchema{
				namespaceLabel: namespaceFieldSchema,
			},
		}
		expectedErr := fmt.Errorf("error")

		createAccountUC.EXPECT().WithStorage(storage).Return(createAccountUC)
		createAccountUC.EXPECT().Execute(ctx, "myNamespace", "").Return(nil, expectedErr)

		response, err := operation.Handler()(ctx, request, data)

		assert.Empty(t, response)
		assert.Equal(t, expectedErr, err)
	})
}
