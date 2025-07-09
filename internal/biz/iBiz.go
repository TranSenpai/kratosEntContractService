package biz

import (
	"context"
	models "kratosEntContractService/internal/models"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewContractService)

type IContractBiz interface {
	CreateContract(ctx context.Context, contract *models.CreateContract) error
	UpdateContract(ctx context.Context, contract *models.UpdateContract, filter *models.ContractFilter) error
	DeleteContract(ctx context.Context, filter *models.ContractFilter) error
	GetContract(ctx context.Context, filter *models.ContractFilter) (*models.ReplyContract, error)
	ListContract(ctx context.Context, filter *models.ContractFilter) ([]models.ReplyContract, error)
	ListTotalContractEachRoom(ctx context.Context, number uint8) ([]models.TotalContractsEachRoom, error)
	SignContract(ctx context.Context, filter *models.ContractFilter, signature string) error
}

type contractBiz struct {
	contractRepo IContractRepo
}

var ContractBiz *contractBiz

func NewContractService(contracRepo IContractRepo) IContractBiz {
	if ContractBiz == nil {
		ContractBiz = &contractBiz{contractRepo: contracRepo}
	}
	return ContractBiz
}
