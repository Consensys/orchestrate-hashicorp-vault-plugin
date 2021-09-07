package keys

import (
	"fmt"
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

func (s *keysCtrlTestSuite) TestKeysController_Get() {
	path := s.controller.Paths()[2]
	getOperation := path.Operations[logical.ReadOperation]

	s.T().Run("should define the correct path", func(t *testing.T) {
		assert.Equal(t, fmt.Sprintf("keys/%s", framework.GenericNameRegex(formatters.IDLabel)), path.Pattern)
		assert.NotEmpty(t, getOperation)
	})

	s.T().Run("should define correct properties", func(t *testing.T) {
		properties := getOperation.Properties()

		assert.NotEmpty(t, properties.Description)
		assert.NotEmpty(t, properties.Summary)
		assert.NotEmpty(t, properties.Examples[0].Description)
		assert.NotEmpty(t, properties.Examples[0].Data)
		assert.NotEmpty(t, properties.Examples[0].Response)
		assert.NotEmpty(t, properties.Responses[200])
		assert.NotEmpty(t, properties.Responses[400])
		assert.NotEmpty(t, properties.Responses[500])
	})

	s.T().Run("handler should execute the correct use case", func(t *testing.T) {
		key := apputils.FakeKey()
		request := &logical.Request{
			Storage: s.storage,
			Headers: map[string][]string{
				formatters.NamespaceHeader: {key.Namespace},
			},
		}
		data := &framework.FieldData{
			Raw: map[string]interface{}{
				formatters.IDLabel: key.ID,
			},
			Schema: map[string]*framework.FieldSchema{
				formatters.IDLabel: formatters.IDFieldSchema,
			},
		}

		s.getKeyUC.EXPECT().Execute(gomock.Any(), key.ID, key.Namespace).Return(key, nil)

		response, err := getOperation.Handler()(s.ctx, request, data)
		require.NoError(t, err)

		assert.Equal(t, key.PublicKey, response.Data[formatters.PublicKeyLabel])
		assert.Equal(t, key.Namespace, response.Data[formatters.NamespaceLabel])
		assert.Equal(t, key.Algorithm, response.Data[formatters.AlgorithmLabel])
		assert.Equal(t, key.Curve, response.Data[formatters.CurveLabel])
		assert.Equal(t, key.ID, response.Data[formatters.IDLabel])
		assert.Equal(t, key.Tags, response.Data[formatters.TagsLabel])
		assert.NotNil(t, key.CreatedAt, response.Data[formatters.CreatedAtLabel])
		assert.NotNil(t, key.UpdatedAt, response.Data[formatters.UpdatedAtLabel])
		assert.Equal(t, 1, response.Data[formatters.VersionLabel])
	})

	s.T().Run("should map errors correctly and return the correct http status", func(t *testing.T) {
		request := &logical.Request{
			Storage: s.storage,
		}
		data := &framework.FieldData{
			Raw: map[string]interface{}{
				formatters.IDLabel: "my-key",
			},
			Schema: map[string]*framework.FieldSchema{
				formatters.IDLabel: formatters.IDFieldSchema,
			},
		}
		expectedErr := errors.NotFoundError("error")

		s.getKeyUC.EXPECT().Execute(gomock.Any(), "my-key", "").Return(nil, expectedErr)

		_, err := getOperation.Handler()(s.ctx, request, data)

		assert.Equal(t, err, logical.ErrUnsupportedPath)
	})
}
