package ethereum

import (
	ethereum "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/ethereum/use-cases"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/ethereum/use-cases/mocks"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type ethereumCtrlTestSuite struct {
	suite.Suite
	createAccountUC *mocks.MockCreateAccountUseCase
	controller      *controller
}

func (s *ethereumCtrlTestSuite) CreateAccount() ethereum.CreateAccountUseCase {
	return s.createAccountUC
}

func (s *ethereumCtrlTestSuite) SignPayload() ethereum.SignUseCase {
	return nil
}

func (s *ethereumCtrlTestSuite) SignTransaction() ethereum.SignTransactionUseCase {
	return nil
}

func (s *ethereumCtrlTestSuite) SignQuorumPrivateTransaction() ethereum.SignQuorumPrivateTransactionUseCase {
	return nil
}

func (s *ethereumCtrlTestSuite) SignEEATransaction() ethereum.SignEEATransactionUseCase {
	return nil
}

var _ ethereum.UseCases = &ethereumCtrlTestSuite{}

func TestEthereumController(t *testing.T) {
	s := new(ethereumCtrlTestSuite)
	suite.Run(t, s)
}

func (s *ethereumCtrlTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	s.createAccountUC = mocks.NewMockCreateAccountUseCase(ctrl)
	s.controller = NewController(s)
}

func (s *ethereumCtrlTestSuite) TestEthereumController_Paths() {
	s.T().Run("should define the correct paths", func(t *testing.T) {
		paths := s.controller.Paths()

		assert.Equal(t, "ethereum/accounts", paths[0].Pattern)
		assert.NotEmpty(t, paths[0].Operations[logical.CreateOperation])

		assert.Equal(t, "ethereum/accounts/import", paths[1].Pattern)
		assert.NotEmpty(t, paths[1].Operations[logical.CreateOperation])
	})
}
