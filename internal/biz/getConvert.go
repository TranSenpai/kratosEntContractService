package biz

import (
	entity "dormitory/internal/entities"
	models "dormitory/internal/models"
	"errors"
	"net/http"
)

func (c contractBiz) convertCreateDto(contract *models.CreateContract) (*entity.Contract, error) {
	dtoAdapter := NewDtoApdater()
	dtoAdapter.SetDtoAdapter("createDto", NewCreateDto(contract))
	createDto := dtoAdapter.GetDtoAdapter("createDto")
	err := createDto.Convert()
	if err != nil {
		return nil, err
	}
	data := createDto.Getter()
	if data == nil {
		return nil, GetError(http.StatusInternalServerError, errors.New("failed to assertion create DTO"))
	}
	entity, ok := data.(*entity.Contract)
	if !ok {
		return nil, GetError(http.StatusInternalServerError, errors.New("failed to assertion create DTO"))
	}
	return entity, nil
}

func (c contractBiz) convertUpdateDto(contract *models.UpdateContract) (map[string]any, error) {
	dtoAdapter := NewDtoApdater()
	dtoAdapter.SetDtoAdapter("updateDto", NewUpdateDto(contract))
	createDto := dtoAdapter.GetDtoAdapter("updateDto")
	err := createDto.Convert()
	if err != nil {
		return nil, err
	}
	data := createDto.Getter()
	mapField, ok := data.(map[string]any)
	if !ok {
		return nil, GetError(http.StatusInternalServerError, errors.New("faile to assertion update DTO"))
	}
	return mapField, nil
}

func (c contractBiz) convertGetListDto(entity *entity.Contract) (*models.ReplyContract, error) {
	dtoAdapter := NewDtoApdater()
	dtoAdapter.SetDtoAdapter("getListDto", NewGetListDto(entity))
	getListDto := dtoAdapter.GetDtoAdapter("getListDto")
	err := getListDto.Convert()
	if err != nil {
		return nil, err
	}
	data := getListDto.Getter()
	contract, ok := data.(*models.ReplyContract)
	if !ok {
		return nil, GetError(http.StatusInternalServerError, errors.New("faile to assertion get list DTO"))
	}

	return contract, nil
}
