package data

import (
	models "dormitory/internal/models"

	"gorm.io/gorm"
)

type query struct {
	tx *gorm.DB
}

func buildQueryIncludeExclude[T any](slice []T, q *query, field string) *query {
	if len(slice) == 0 {
		return q
	}
	q.tx = q.tx.Where(field+" IN ?", slice)

	return q
}

func (q *query) studentInfoInclude(filter *models.ContractFilter) *query {
	if filter == nil {
		return nil
	}
	q = buildQueryIncludeExclude(filter.Id.Includes, q, "id")
	q = buildQueryIncludeExclude(filter.StudentCode.Includes, q, "student_code")
	q = buildQueryIncludeExclude(filter.Email.Includes, q, "email")
	q = buildQueryIncludeExclude(filter.FirstName.Includes, q, "first_name")
	q = buildQueryIncludeExclude(filter.LastName.Includes, q, "last_name")
	q = buildQueryIncludeExclude(filter.MiddleName.Includes, q, "middle_name")

	return q

}

func (q *query) contractInfoInclude(filter *models.ContractFilter) *query {
	if filter == nil {
		return nil
	}
	q = buildQueryIncludeExclude(filter.Phone.Includes, q, "phone")
	q = buildQueryIncludeExclude(filter.Sign.Includes, q, "sign")
	q = buildQueryIncludeExclude(filter.RoomId.Includes, q, "room_id")
	q = buildQueryIncludeExclude(filter.Gender.Includes, q, "gender")
	q = buildQueryIncludeExclude(filter.Address.Includes, q, "address")

	return q
}

func (q *query) studentInfoExclude(filter *models.ContractFilter) *query {
	if filter == nil {
		return nil
	}
	q = buildQueryIncludeExclude(filter.Id.Excludes, q, "id")
	q = buildQueryIncludeExclude(filter.StudentCode.Excludes, q, "student_code")
	q = buildQueryIncludeExclude(filter.Email.Excludes, q, "email")
	q = buildQueryIncludeExclude(filter.FirstName.Excludes, q, "first_name")
	q = buildQueryIncludeExclude(filter.LastName.Excludes, q, "last_name")
	q = buildQueryIncludeExclude(filter.MiddleName.Excludes, q, "middle_name")

	return q
}

func (q *query) contractInfoExclude(filter *models.ContractFilter) *query {
	if filter == nil {
		return nil
	}
	q = buildQueryIncludeExclude(filter.Phone.Excludes, q, "phone")
	q = buildQueryIncludeExclude(filter.Sign.Excludes, q, "sign")
	q = buildQueryIncludeExclude(filter.RoomId.Excludes, q, "room_id")
	q = buildQueryIncludeExclude(filter.Gender.Excludes, q, "gender")
	q = buildQueryIncludeExclude(filter.Address.Excludes, q, "address")

	return q
}

func (q *query) buildIncludeQuery(filter *models.ContractFilter) *query {
	if filter == nil {
		return nil
	}

	return q.studentInfoInclude(filter).contractInfoInclude(filter)
}

func (q *query) buildExcludeQuery(filter *models.ContractFilter) *query {
	if filter == nil {
		return nil
	}

	return q.studentInfoExclude(filter).contractInfoExclude(filter)
}

func (q *query) buildQuery(filter *models.ContractFilter) *gorm.DB {
	q.buildIncludeQuery(filter).buildExcludeQuery(filter)
	if filter.IsActive != nil {
		q.tx.Where("is_active = ?", *filter.IsActive)
	}
	if filter.RegistryAt.FromTime != nil && !filter.RegistryAt.FromTime.IsZero() {
		q.tx = q.tx.Where("registry_at >= ?", *filter.RegistryAt.FromTime)
	}
	if filter.RegistryAt.ToTime != nil && !filter.RegistryAt.ToTime.IsZero() {
		q.tx = q.tx.Where("registry_at <= ?", *filter.RegistryAt.ToTime)
	}

	return q.tx
}

func NewQueryBuilder(tx *gorm.DB) *query {
	return &query{
		tx: tx,
	}
}
