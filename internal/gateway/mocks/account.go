package mocks

import (
	"github.com/Daniel-Vinicius/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/mock"
)

type AccountGatewayMock struct {
	mock.Mock
}

func NewAccountGatewayMock() *AccountGatewayMock {
	return &AccountGatewayMock{}
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}
