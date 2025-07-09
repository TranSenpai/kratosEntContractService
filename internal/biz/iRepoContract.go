package biz

import (
	"context"
	entity "dormitory/internal/ent"
	models "dormitory/internal/models"
)

type IContractRepo interface {
	CreateContract(ctx context.Context, contract *entity.Contract) error
	UpdateContract(ctx context.Context, contract map[string]any, filter *models.ContractFilter) error
	DeleteContract(ctx context.Context, filter *models.ContractFilter) error
	GetContract(ctx context.Context, filter *models.ContractFilter) (*entity.Contract, error)
	ListContract(ctx context.Context, filter *models.ContractFilter) ([]entity.Contract, error)
	GetTotalContractRoom(ctx context.Context, roomID string) (models.TotalContractsEachRoom, error)
	ListTotalContractEachRoom(ctx context.Context, number uint8) ([]models.TotalContractsEachRoom, error)
}
