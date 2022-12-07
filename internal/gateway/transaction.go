package gateway

import "github.com/Daniel-Vinicius/fc-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
