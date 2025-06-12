package data

import (
	models "dormitory/internal/models"

	"gorm.io/gorm"
)

func getQuery[T any](data models.SearchFilter[T], tx *gorm.DB, field string) *gorm.DB {
	if len(data.Includes) > 0 {
		tx = tx.Where(field+" IN ?", data.Includes)
	}
	if len(data.Excludes) > 0 {
		tx = tx.Where(field+" NOT IN ?", data.Excludes)
	}

	return tx
}

func (cr contractRepo) buildQuery(filter *models.ContractFilter, tx *gorm.DB) *gorm.DB {
	if filter == nil {
		return nil
	}
	tx = getQuery(filter.Id, tx, "id")
	tx = getQuery(filter.StudentCode, tx, "student_code")
	tx = getQuery(filter.Email, tx, "email")
	tx = getQuery(filter.FirstName, tx, "first_name")
	tx = getQuery(filter.LastName, tx, "last_name")
	tx = getQuery(filter.MiddleName, tx, "middle_name")
	tx = getQuery(filter.Phone, tx, "phone")
	tx = getQuery(filter.Sign, tx, "sign")
	tx = getQuery(filter.RoomId, tx, "room_id")
	tx = getQuery(filter.Gender, tx, "gender")
	tx = getQuery(filter.Address, tx, "address")
	// tx = nil
	return tx
}
