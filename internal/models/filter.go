package requestModel

import "time"

type SearchFilter[T any] struct {
	Includes []T
	Excludes []T
}

type RequestTime struct {
	FromTime *time.Time
	ToTime   *time.Time
}

type ContractFilter struct {
	Id          SearchFilter[uint64]
	StudentCode SearchFilter[string]
	FirstName   SearchFilter[string]
	LastName    SearchFilter[string]
	MiddleName  SearchFilter[string]
	Email       SearchFilter[string]
	Sign        SearchFilter[string]
	Phone       SearchFilter[string]
	Gender      SearchFilter[uint32]
	Address     SearchFilter[string]
	RoomId      SearchFilter[string]
	IsActive    *bool
	RegistryAt  RequestTime
}
