package createclient

import (
	"time"

	"github.com/Daniel-Vinicius/fc-ms-wallet/internal/entity"
	"github.com/Daniel-Vinicius/fc-ms-wallet/internal/gateway"
)

type CreateClientInputDTO struct {
	Name string
	Email string
}

type CreateClientOutputDTO struct {
	ID string
	Name string
	Email string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (useCase *CreateClientUseCase) Execute(input *CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = useCase.ClientGateway.Save(client)
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		ID: client.ID,
		Name: client.Name,
		Email: client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil
}