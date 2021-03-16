package keys

import (
	"context"
	"fmt"
	"testing"

	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/log"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateKey_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)
	ctx := log.Context(context.Background(), log.Default())

	usecase := NewCreateKeyUseCase().WithStorage(mockStorage)

	t.Run("should execute use case successfully by generating a private key", func(t *testing.T) {
		fakeKey := utils.FakeKey()
		mockStorage.EXPECT().Put(ctx, gomock.Any()).Return(nil)

		key, err := usecase.Execute(ctx, fakeKey.Namespace, fakeKey.ID, fakeKey.Algorithm, fakeKey.Curve, "", fakeKey.Tags)

		assert.NoError(t, err)
		assert.Equal(t, fakeKey.Namespace, key.Namespace)
		assert.NotEmpty(t, key.PublicKey)
	})

	t.Run("should fail with same error if Put fails", func(t *testing.T) {
		expectedErr := fmt.Errorf("error")

		mockStorage.EXPECT().Put(ctx, gomock.Any()).Return(expectedErr)

		key, err := usecase.Execute(ctx, "namespace", "id", "algo", "curve", "", map[string]string{})
		assert.Nil(t, key)
		assert.Equal(t, expectedErr, err)
	})
}
