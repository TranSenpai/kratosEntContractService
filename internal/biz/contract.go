package biz

import (
	"context"
	entity "dormitory/internal/entities"
	models "dormitory/internal/models"
	"errors"
	"net/http"
)

func (c *contractBiz) CreateContract(ctx context.Context, contract *models.CreateContract) error {
	if contract == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("contract empty"))
	}
	checker := c.GetCheckRequiredField()
	if err := checker.CheckRequiredField(ctx, contract); err != nil {
		return err
	}
	entity, err := c.convertCreateDto(contract)
	if err != nil {
		return err
	}

	return c.contractRepo.CreateContract(ctx, entity)
}

func (c *contractBiz) UpdateContract(ctx context.Context, contract *models.UpdateContract, filter *models.ContractFilter) error {
	if contract == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("nil contract"))
	}
	mapField, err := c.convertUpdateDto(contract)
	if err != nil {
		return err
	}

	return c.contractRepo.UpdateContract(ctx, mapField, filter)
}

func (c *contractBiz) DeleteContract(ctx context.Context, contractID uint64) error {
	contract, err := c.contractRepo.GetContract(ctx, contractID)
	if err != nil {
		return err
	}
	if contract.IsActive {
		return GetError(http.StatusBadRequest, errors.New("contract is available, terminate the contract before delete"))
	}

	return c.contractRepo.DeleteContract(ctx, contractID)
}

func (c *contractBiz) GetContract(ctx context.Context, contractID uint64) (*models.ReplyContract, error) {
	contractEntity, err := c.contractRepo.GetContract(ctx, contractID)
	if err != nil {
		return nil, err
	}
	if contractEntity.ID == 0 {
		return nil, GetError(http.StatusNotFound, errors.New("contract not found"))
	}
	contract, err := c.convertGetListDto(contractEntity)
	if err != nil {
		return nil, err
	}

	return contract, nil
}

func (c *contractBiz) ListContract(ctx context.Context, filter *models.ContractFilter) ([]models.ReplyContract, error) {
	entities, err := c.contractRepo.ListContract(ctx, filter)
	if err != nil {
		return nil, err
	}
	var contracts []models.ReplyContract
	for _, v := range entities {
		data, err := c.convertGetListDto(&v)
		if err != nil || data == nil {
			return nil, err
		}
		contracts = append(contracts, *data)
	}

	return contracts, nil
}

func (c *contractBiz) ListTotalContractEachRoom(ctx context.Context, number uint8) ([]models.TotalContractsEachRoom, error) {
	totalContract, err := c.contractRepo.ListTotalContractEachRoom(ctx, number)
	if err != nil {
		return nil, err
	}

	return totalContract, nil
}

func (c *contractBiz) ValidateSignContract(ctx context.Context, contract *entity.Contract, signature string) error {
	if contract == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("nil contract"))
	}
	if contract.IsActive {
		return GetError(http.StatusUnprocessableEntity, errors.New("contract is active"))
	}
	if contract.Sign != "" {
		return GetError(http.StatusUnprocessableEntity, errors.New("contract is active"))
	}
	contract.IsActive = true
	contract.Sign = signature

	return nil
}

func (c *contractBiz) SignContract(ctx context.Context, contractID uint64, signature string) error {
	contractEntity, err := c.contractRepo.GetContract(ctx, contractID)
	if err != nil {
		return err
	}
	err = c.ValidateSignContract(ctx, contractEntity, signature)
	if err != nil {
		return err
	}
	var mapField = map[string]any{
		"ID":       contractEntity.ID,
		"IsActive": true,
		"Sign":     signature,
	}

	var filter models.ContractFilter
	filter.Id.Includes = append(filter.Id.Includes, contractID)
	err = c.contractRepo.UpdateContract(ctx, mapField, &filter)
	if err != nil {
		return err
	}

	return nil
}
