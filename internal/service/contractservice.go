package service

import (
	"context"
	contractApi "dormitory/api/contract"
	"dormitory/internal/biz"
	models "dormitory/internal/models"
	"time"
	// "github.com/go-kratos/kratos/v2/api/metadata"
)

type ContractService struct {
	contractApi.UnimplementedContractServiceServer
	bizContract biz.IContractBiz
}

func NewContractService(b biz.IContractBiz) *ContractService {
	return &ContractService{
		bizContract: b,
	}
}

// protobuf convert timestamp to UTC
func (s *ContractService) CreateContract(ctx context.Context, req *contractApi.CreateContractRequest) (*contractApi.CreateContractReply, error) {
	contract := s.ConvertCreateContractRequest(req)
	// contract.RegistryAt = time.Date(2025, time.June, 17, 19, 30, 0, 0, time.UTC)
	contract.RegistryAt = time.Now()
	err := s.bizContract.CreateContract(ctx, contract)
	if err != nil {
		return nil, err
	}

	return &contractApi.CreateContractReply{Message: "Create successfully"}, nil
}

func (s *ContractService) UpdateContract(ctx context.Context, req *contractApi.UpdateContractRequest) (*contractApi.UpdateContractReply, error) {
	contract := s.ConvertUpdateContractRequest(req)
	var filter models.ContractFilter
	filter.Id.Includes = append(filter.Id.Includes, req.Id)
	if err := s.bizContract.UpdateContract(ctx, contract, &filter); err != nil {
		return nil, err
	}
	return &contractApi.UpdateContractReply{
		Message: "Successfully",
	}, nil
}

func (s *ContractService) SignContract(ctx context.Context, req *contractApi.SignRequest) (*contractApi.SignReply, error) {
	err := s.bizContract.SignContract(ctx, req.Id, req.Sign)
	if err != nil {
		return nil, err
	}
	return &contractApi.SignReply{
		Message: "Sign successfully",
	}, nil
}

func (s *ContractService) DeleteContract(ctx context.Context, req *contractApi.DeleteContractRequest) (*contractApi.DeleteContractReply, error) {
	contractID := req.Id
	if err := s.bizContract.DeleteContract(ctx, contractID); err != nil {
		return nil, err
	}
	return &contractApi.DeleteContractReply{Message: "Successfully"}, nil
}

func (s *ContractService) GetContract(ctx context.Context, req *contractApi.GetContractRequest) (*contractApi.GetContractReply, error) {
	contract, err := s.bizContract.GetContract(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	result := s.ConvertToContractReply(contract)
	return &contractApi.GetContractReply{Contract: result, Message: "Successfully"}, nil
}

func (s *ContractService) ListContract(ctx context.Context, req *contractApi.ListContractRequest) (*contractApi.ListContractReply, error) {
	filter, err := s.ConvertFilter(req)
	if err != nil {
		return nil, err
	}
	contracts, err := s.bizContract.ListContract(ctx, filter)
	if err != nil {
		return nil, err
	}
	var result contractApi.ListContractReply
	for _, v := range contracts {
		result.Contract = append(result.Contract, s.ConvertToContractReply(&v))
	}

	return &result, nil
}

func (s *ContractService) ListTotalContractEachRoom(ctx context.Context, req *contractApi.ListTotalContractEachRoomRequest) (*contractApi.ListTotalContractEachRoomReply, error) {
	totalContracts, err := s.bizContract.ListTotalContractEachRoom(ctx, uint8(req.Number))
	if err != nil {
		return nil, err
	}
	var totalContractsReply contractApi.ListTotalContractEachRoomReply
	for _, v := range totalContracts {
		var totalContractOneRoom contractApi.ListTotalContractEachRoomReply_ContractEachRoom
		totalContractOneRoom.RoomId = v.RoomID
		totalContractOneRoom.Total = uint32(v.Total)
		totalContractsReply.ContractEachRoom = append(totalContractsReply.ContractEachRoom, &totalContractOneRoom)
	}

	return &totalContractsReply, nil
}
