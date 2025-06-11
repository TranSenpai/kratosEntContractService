package data

import (
	"context"
	"dormitory/internal/biz"
	entity "dormitory/internal/entities"
	models "dormitory/internal/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// Type conversion the value nil to type concrete (*contractRepo)
// After that we assign some variable to the pointer type (var x Interface = (*ConcreteInterface)(nil))
// Compiler will check if ConcreteInterface satisfy of not (implement all method of interface)
// Use to check if a concrete type fully implement it's Interface
// In this example b/c we dont use the value so let it be a blank variable:
var _ biz.IContractRepo = (*contractRepo)(nil)

func NewContractRepo(dbConnection *gorm.DB) biz.IContractRepo {
	return &contractRepo{
		db: dbConnection,
	}
}

type contractRepo struct {
	db *gorm.DB
}

func (cr *contractRepo) CreateContract(ctx context.Context, createContract *entity.Contract) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		// str := gorm.ErrInvalidTransaction.Error()
		// return GetError(errors.New(str))
		err := tx.Debug().Create(createContract).WithContext(ctx).Error
		if err != nil {
			return GetError(err)
		}
		return nil
	})
}

func (cr *contractRepo) UpdateContract(ctx context.Context, updateContract map[string]any, filter *models.ContractFilter) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		// Updates supports updating with struct or map[string]interface{},
		// when updating with struct it will only update non-zero fields by default
		q := NewQueryBuilder(tx)
		queryTx := q.buildQuery(filter)
		result := queryTx.Debug().Model(&entity.Contract{}).Omit("id").Updates(updateContract).WithContext(ctx)
		if result.Error != nil {
			return GetError(result.Error)
		}

		return nil
	})
}

func (cr *contractRepo) DeleteContract(ctx context.Context, contractID uint64) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Debug().Where("id = ?", contractID).Delete(&entity.Contract{}).WithContext(ctx).Error
		if err != nil {
			return GetError(errors.New(gorm.ErrRecordNotFound.Error()))
		}

		return nil
	})
}

func (cr contractRepo) GetContract(ctx context.Context, contractID uint64) (*entity.Contract, error) {
	var contract entity.Contract
	err := cr.db.Debug().Model(&entity.Contract{}).Where("id = ?", contractID).Find(&contract).Error
	if err != nil {
		return &entity.Contract{}, GetError(errors.New(gorm.ErrRecordNotFound.Error()))
	}

	return &contract, nil
}

func (cr contractRepo) ListContract(ctx context.Context, filter *models.ContractFilter) ([]entity.Contract, error) {
	var lst []entity.Contract
	q := NewQueryBuilder(cr.db)
	cr.db = q.buildQuery(filter)

	// err := cr.db.Debug().Model(&entity.Contract{}).Find(&lst).Error
	partitionStr := CallPartition2025(filter)
	fmt.Println(partitionStr)
	err := cr.db.Debug().Table(partitionStr).Find(&lst).Error

	if err != nil {
		GetError(err)
	}

	return lst, err
}

func (cr contractRepo) GetTotalContractRoom(ctx context.Context, roomID string) (models.TotalContractsEachRoom, error) {
	var result models.TotalContractsEachRoom
	err := cr.db.Debug().Model(&entity.Contract{}).
		Select("COUNT(id) as total, room_id").Where("is_active = ? and room_id = ?", true, roomID).
		Group("room_id").Having("total > ?", 4).Find(&result).Error
	if err != nil {
		GetError(err)
	}

	return result, err
}

func (cr contractRepo) ListTotalContractEachRoom(ctx context.Context, number uint8) ([]models.TotalContractsEachRoom, error) {
	var result []models.TotalContractsEachRoom
	err := cr.db.Debug().Model(&entity.Contract{}).
		Select("COUNT(id) as total, room_id").Group("room_id").
		Having("total >= ?", number).Find(&result).Error
	if err != nil {
		GetError(err)
	}

	return result, err
}
