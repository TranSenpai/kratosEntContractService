package service

import (
	contractApi "dormitory/api/contract"
	models "dormitory/internal/models"
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (s ContractService) ConvertCreateContractRequest(req *contractApi.CreateContractRequest) *models.CreateContract {
	var contract models.CreateContract
	contract.StudentCode = req.StudentCode
	contract.FirstName = req.FirstName
	contract.LastName = req.LastName
	contract.MiddleName = req.MiddleName
	contract.Email = req.Email
	contract.Sign = req.Sign
	contract.Phone = req.Phone
	contract.Gender = req.Gender
	contract.Dob = req.Dob.AsTime()
	contract.Address = req.Address
	contract.Avatar = req.Avatar
	contract.IsActive = req.IsActive
	contract.RoomId = req.RoomId
	contract.NotificationChannels = req.NotificationChannels
	return &contract
}

func (s ContractService) ConvertUpdateContractRequest(req *contractApi.UpdateContractRequest) *models.UpdateContract {
	var contract models.UpdateContract
	contract.StudentCode = req.StudentCode
	contract.FirstName = req.FirstName
	contract.LastName = req.LastName
	contract.MiddleName = req.MiddleName
	contract.Email = req.Email
	contract.Sign = req.Sign
	contract.Phone = req.Phone
	contract.Gender = req.Gender
	time := req.Dob.AsTime()
	contract.Dob = &time
	contract.Address = req.Address
	contract.Avatar = req.Avatar
	contract.IsActive = req.IsActive
	contract.RoomID = req.RoomId
	contract.NotificationChannels = req.NotificationChannels
	return &contract
}

func (s ContractService) ConvertToContractReply(contract *models.ReplyContract, timezone int32) *contractApi.Contract {
	var contractReply contractApi.Contract
	contractReply.Id = contract.Id
	contractReply.StudentCode = contract.StudentCode
	contractReply.FirstName = contract.FirstName
	contractReply.LastName = contract.LastName
	contractReply.MiddleName = contract.MiddleName
	contractReply.Email = contract.Email
	contractReply.Phone = contract.Phone
	contractReply.Gender = contract.Gender
	contractReply.Address = contract.Address
	contractReply.Avatar = contract.Avatar
	contractReply.RoomId = contract.RoomId
	contractReply.IsActive = contract.IsActive
	contractReply.Sign = contract.Sign
	contractReply.NotificationChannels = contract.NotificationChannels
	if contract.RegistryAt != nil {
		contractReply.RegistryAt = timestamppb.New(*contract.RegistryAt)
	}
	if contract.Dob != nil {
		convertTime := contract.Dob.Unix()
		convertTime += int64(timezone * 3600)
		clientTime := time.Unix(convertTime, 0)
		contractReply.Dob = timestamppb.New(clientTime)
	}
	return &contractReply
}
