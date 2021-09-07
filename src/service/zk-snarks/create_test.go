package zksnarks

import (
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/pkg/errors"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/service/formatters"
	apputils "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/stretchr/testify/assert"
)

func (s *zksCtrlTestSuite) TestZksController_Create() {
	path := s.controller.Paths()[0]
	createOperation := path.Operations[logical.CreateOperation]

	s.T().Run("should define the correct path", func(t *testing.T) {
		assert.Equal(t, "zk-snarks/accounts/?", path.Pattern)
		assert.NotEmpty(t, createOperation)
	})

	s.T().Run("should define correct properties", func(t *testing.T) {
		properties := createOperation.Properties()

		assert.NotEmpty(t, properties.Description)
		assert.NotEmpty(t, properties.Summary)
		assert.NotEmpty(t, properties.Examples[0].Description)
		assert.Empty(t, properties.Examples[0].Data)
		assert.NotEmpty(t, properties.Examples[0].Response)
		assert.NotEmpty(t, properties.Responses[200])
		assert.NotEmpty(t, properties.Responses[500])
	})

	s.T().Run("handler should execute the correct use case", func(t *testing.T) {
		account := apputils.FakeZksAccount()
		request := &logical.Request{
			Storage: s.storage,
			Headers: map[string][]string{
				formatters.NamespaceHeader: {account.Namespace},
			},
		}

		s.createAccountUC.EXPECT().Execute(gomock.Any(), account.Namespace).Return(account, nil)

		response, err := createOperation.Handler()(s.ctx, request, &framework.FieldData{})
		require.NoError(t, err)

		assert.Equal(t, account.PublicKey, response.Data[formatters.PublicKeyLabel])
		assert.Equal(t, account.Namespace, response.Data[formatters.NamespaceLabel])
		assert.Equal(t, account.Algorithm, response.Data[formatters.AlgorithmLabel])
		assert.Equal(t, account.Curve, response.Data[formatters.CurveLabel])
	})

	s.T().Run("should map errors correctly and return the correct http status", func(t *testing.T) {
		request := &logical.Request{
			Storage: s.storage,
		}
		expectedErr := errors.NotFoundError("error")

		s.createAccountUC.EXPECT().Execute(gomock.Any(), "").Return(nil, expectedErr)

		_, err := createOperation.Handler()(s.ctx, request, &framework.FieldData{})

		assert.Equal(t, err, logical.ErrUnsupportedPath)
	})
}
