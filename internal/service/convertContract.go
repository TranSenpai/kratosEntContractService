package service

import (
	contractApi "dormitory/api/contract"
	models "dormitory/internal/models"

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
	if req.Dob != nil {
		dob := req.Dob.AsTime()
		contract.Dob = &dob
	}
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

func (s ContractService) ConvertToContractReply(contract *models.ReplyContract) *contractApi.Contract {
	var contractReply contractApi.Contract
	contractReply.Id = int32(contract.Id)
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
	contractReply.RegistryAt = timestamppb.New(contract.RegistryAt)
	contractReply.Dob = timestamppb.New(contract.Dob)

	return &contractReply
}
