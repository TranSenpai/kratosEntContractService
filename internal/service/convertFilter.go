package service

import (
	contractApi "dormitory/api/contract"
	models "dormitory/internal/models"
	"net/http"
	"time"

	kerror "github.com/go-kratos/kratos/v2/errors"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func CheckRequest[T any](slice []T) []T {
	if len(slice) > 0 {
		return slice
	}
	return nil
}

func (s ContractService) ConvertStudentInfo(filter *models.ContractFilter, req *contractApi.ListContractRequest) error {
	if filter == nil {
		return kerror.New(http.StatusUnprocessableEntity, "Service error|", "nil filter")
	}
	if req.Id != nil {
		var idInclude, idExclude []int
		for _, v := range req.Id.Includes {
			idInclude = append(idInclude, int(v))
		}
		for _, v := range req.Id.Excludes {
			idExclude = append(idExclude, int(v))
		}
		filter.Id.Includes = CheckRequest(idInclude)
		filter.Id.Excludes = CheckRequest(idExclude)
	}
	if req.StudentCode != nil {
		filter.StudentCode.Includes = CheckRequest(req.StudentCode.Includes)
		filter.StudentCode.Excludes = CheckRequest(req.StudentCode.Excludes)
	}
	if req.FirstName != nil {
		filter.FirstName.Includes = CheckRequest(req.FirstName.Includes)
		filter.FirstName.Excludes = CheckRequest(req.FirstName.Excludes)
	}
	if req.LastName != nil {
		filter.LastName.Includes = CheckRequest(req.LastName.Includes)
		filter.LastName.Excludes = CheckRequest(req.LastName.Excludes)
	}
	if req.MiddleName != nil {
		filter.MiddleName.Includes = CheckRequest(req.MiddleName.Includes)
		filter.MiddleName.Excludes = CheckRequest(req.MiddleName.Excludes)
	}
	if req.Email != nil {
		filter.Email.Includes = CheckRequest(req.Email.Includes)
		filter.Email.Excludes = CheckRequest(req.Email.Excludes)
	}
	if req.Gender != nil {
		filter.Gender.Includes = CheckRequest(req.Gender.Includes)
		filter.Gender.Excludes = CheckRequest(req.Gender.Excludes)
	}
	return nil
}

func (s ContractService) partRequestRegistry(time *timestamppb.Timestamp) *time.Time {
	if time != nil {
		fromTime := time.AsTime()
		return &fromTime
	}

	return nil
}

func (s ContractService) ConvertContractInfo(filter *models.ContractFilter, req *contractApi.ListContractRequest) error {
	if filter == nil {
		return kerror.New(http.StatusUnprocessableEntity, "Service error|", "nil filter")
	}
	if req.Sign != nil {
		filter.Sign.Includes = CheckRequest(req.Sign.Includes)
		filter.Sign.Excludes = CheckRequest(req.Sign.Excludes)
	}
	if req.Phone != nil {
		filter.Phone.Includes = CheckRequest(req.Phone.Includes)
		filter.Phone.Excludes = CheckRequest(req.Phone.Excludes)
	}
	if req.Address != nil {
		filter.Address.Includes = CheckRequest(req.Address.Includes)
		filter.Address.Excludes = CheckRequest(req.Address.Excludes)
	}
	if req.RoomId != nil {
		filter.RoomId.Includes = CheckRequest(req.RoomId.Includes)
		filter.RoomId.Excludes = CheckRequest(req.RoomId.Excludes)
	}
	filter.IsActive = req.IsActive
	if req.RegistryAt != nil {
		filter.RegistryAt.FromTime = s.partRequestRegistry(req.RegistryAt.FromTime)
		filter.RegistryAt.ToTime = s.partRequestRegistry(req.RegistryAt.ToTime)
	}

	return nil
}

func (s ContractService) ConvertFilter(req *contractApi.ListContractRequest) (*models.ContractFilter, error) {
	var filter models.ContractFilter
	err := s.ConvertStudentInfo(&filter, req)
	if err != nil {
		return nil, err
	}
	err = s.ConvertContractInfo(&filter, req)
	if err != nil {
		return nil, err
	}
	return &filter, nil
}
