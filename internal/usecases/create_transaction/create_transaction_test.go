package createtransaction

import (
	"testing"

	"github.com/Daniel-Vinicius/fc-ms-wallet/internal/entity"
	"github.com/Daniel-Vinicius/fc-ms-wallet/internal/gateway/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "john@doe.com")
	account1 := entity.NewAccount(client1)
	client2, _ := entity.NewClient("John Doe 2", "john2@doe.com")
	account2 := entity.NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	mockAccount := mocks.NewAccountGatewayMock()
	mockAccount.On("FindByID", account1.ID).Return(account1, nil)
	mockAccount.On("FindByID", account2.ID).Return(account2, nil)

	mockTransaction := mocks.NewTransactionGatewayMock()
	mockTransaction.On("Create", mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo: account2.ID,
		Amount: 100,
	}

	useCase := NewCreateTransactionUseCase(mockTransaction, mockAccount)
	output, err := useCase.Execute(inputDto)
	
	assert.Nil(t, err)
	assert.NotNil(t, output)

	mockAccount.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "FindByID", 2)

	mockTransaction.AssertExpectations(t)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)
}