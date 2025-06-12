package data

import (
	models "dormitory/internal/models"

	"gorm.io/gorm"
)

func getQueryInclude[T any](slice []T, tx *gorm.DB, field string) *gorm.DB {
	if len(slice) == 0 {
		return tx
	}

	return tx.Where(field+" IN ?", slice)
}

func getQueryExclude[T any](slice []T, tx *gorm.DB, field string) *gorm.DB {
	if len(slice) == 0 {
		return tx
	}

	return tx.Where(field+" NOT IN ?", slice)
}

func (cr contractRepo) studentInfoInclude(filter *models.ContractFilter, tx *gorm.DB) *gorm.DB {
	if filter == nil {
		return nil
	}
	tx = getQueryInclude(filter.Id.Includes, tx, "id")
	tx = getQueryInclude(filter.StudentCode.Includes, tx, "student_code")
	tx = getQueryInclude(filter.Email.Includes, tx, "email")
	tx = getQueryInclude(filter.FirstName.Includes, tx, "first_name")
	tx = getQueryInclude(filter.LastName.Includes, tx, "last_name")
	tx = getQueryInclude(filter.MiddleName.Includes, tx, "middle_name")

	return tx
}

func (cr contractRepo) contractInfoInclude(filter *models.ContractFilter, tx *gorm.DB) *gorm.DB {
	if filter == nil {
		return nil
	}
	tx = getQueryInclude(filter.Phone.Includes, tx, "phone")
	tx = getQueryInclude(filter.Sign.Includes, tx, "sign")
	tx = getQueryInclude(filter.RoomId.Includes, tx, "room_id")
	tx = getQueryInclude(filter.Gender.Includes, tx, "gender")
	tx = getQueryInclude(filter.Address.Includes, tx, "address")

	return tx
}

func (cr contractRepo) studentInfoExclude(filter *models.ContractFilter, tx *gorm.DB) *gorm.DB {
	if filter == nil {
		return nil
	}
	tx = getQueryExclude(filter.Id.Excludes, tx, "id")
	tx = getQueryExclude(filter.StudentCode.Excludes, tx, "student_code")
	tx = getQueryExclude(filter.Email.Excludes, tx, "email")
	tx = getQueryExclude(filter.FirstName.Excludes, tx, "first_name")
	tx = getQueryExclude(filter.LastName.Excludes, tx, "last_name")
	tx = getQueryExclude(filter.MiddleName.Excludes, tx, "middle_name")

	return tx
}

func (cr contractRepo) contractInfoExclude(filter *models.ContractFilter, tx *gorm.DB) *gorm.DB {
	if filter == nil {
		return nil
	}
	tx = getQueryExclude(filter.Phone.Excludes, tx, "phone")
	tx = getQueryExclude(filter.Sign.Excludes, tx, "sign")
	tx = getQueryExclude(filter.RoomId.Excludes, tx, "room_id")
	tx = getQueryExclude(filter.Gender.Excludes, tx, "gender")
	tx = getQueryExclude(filter.Address.Excludes, tx, "address")

	return tx
}

func (cr contractRepo) buildIncludeQuery(filter *models.ContractFilter, tx *gorm.DB) *gorm.DB {
	if filter == nil {
		return nil
	}
	tx = cr.studentInfoInclude(filter, tx)
	tx = cr.contractInfoInclude(filter, tx)

	return tx
}

func (cr contractRepo) buildExcludeQuery(filter *models.ContractFilter, tx *gorm.DB) *gorm.DB {
	if filter == nil {
		return nil
	}
	tx = cr.studentInfoExclude(filter, tx)
	tx = cr.contractInfoExclude(filter, tx)

	return tx
}

func (cr contractRepo) buildQuery(filter *models.ContractFilter, tx *gorm.DB) *gorm.DB {
	tx = cr.buildIncludeQuery(filter, tx)
	tx = cr.buildExcludeQuery(filter, tx)
	if filter.IsActive != nil {
		tx = tx.Where("is_active = ?", *filter.IsActive)
	}
	if filter.RegistryAt.FromTime != nil && !filter.RegistryAt.FromTime.IsZero() {
		tx = tx.Where("registry_at >= ?", *filter.RegistryAt.FromTime)
	}
	if filter.RegistryAt.ToTime != nil && !filter.RegistryAt.ToTime.IsZero() {
		tx = tx.Where("registry_at <= ?", *filter.RegistryAt.ToTime)
	}

	return tx
}
