package biz

import (
	entity "dormitory/internal/entities"
	models "dormitory/internal/models"
	"encoding/base64"
	"errors"
	"net/http"
)

type IDto interface {
	Convert() error
	Getter() any
}

// factory
type DtoAdapter struct {
	converter map[string]IDto
}

func (d *DtoAdapter) SetDtoAdapter(kind string, dto IDto) {
	d.converter[kind] = dto
}

func (d *DtoAdapter) GetDtoAdapter(kind string) IDto {
	return d.converter[kind]
}

func NewDtoApdater() *DtoAdapter {
	return &DtoAdapter{}
}

func DecodeBase64(input string) (*[]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, GetError(http.StatusInternalServerError, err)
	}

	return &decoded, nil
}

func EncodeBase64(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func decodeAvatar(avatar *string) ([]byte, error) {
	if avatar != nil && *avatar != "" {
		avatarDecoded, err := DecodeBase64(*avatar)
		if err != nil {
			return nil, GetError(http.StatusUnprocessableEntity, err)
		}
		return *avatarDecoded, nil
	}

	return nil, nil
}

func encodeAvatar(avatar []byte) string {
	if avatar != nil {
		avatarDecoded := EncodeBase64(avatar)

		return avatarDecoded
	}

	return ""
}

// concrete Create DTO
type createDto struct {
	contract *models.CreateContract
	entity   *entity.Contract
}

func (createDto *createDto) convertStudentInfo() error {
	if createDto.entity == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("nil entity"))
	}
	createDto.entity.StudentCode = createDto.contract.StudentCode
	createDto.entity.FirstName = createDto.contract.FirstName
	createDto.entity.LastName = createDto.contract.LastName
	createDto.entity.MiddleName = createDto.contract.MiddleName
	createDto.entity.Email = createDto.contract.Email
	createDto.entity.Gender = uint8(createDto.contract.Gender)
	avatarString, err := decodeAvatar(&createDto.contract.Avatar)
	if err != nil {
		return GetError(http.StatusUnprocessableEntity, err)
	}
	createDto.entity.Avatar = avatarString

	return nil
}

func (createDto *createDto) convertContractInfo() {
	createDto.entity.Sign = createDto.contract.Sign
	createDto.entity.Phone = createDto.contract.Phone
	createDto.entity.Address = createDto.contract.Address
	createDto.entity.IsActive = createDto.contract.IsActive
	createDto.entity.RoomID = createDto.contract.RoomId
	createDto.entity.NotificationChannels = uint8(createDto.contract.NotificationChannels)
	createDto.entity.DOB = createDto.contract.Dob
}

func (createDto *createDto) Convert() error {
	err := createDto.convertStudentInfo()
	if err != nil {
		return err
	}
	createDto.convertContractInfo()

	return nil
}

func (createDto *createDto) Getter() any {
	return createDto
}

func NewCreateDto(contract *models.CreateContract) *createDto {
	return &createDto{contract: contract}
}

func (c contractBiz) convertCreateDto(contract *models.CreateContract) (*entity.Contract, error) {
	dtoAdapter := NewDtoApdater()
	dtoAdapter.SetDtoAdapter("createDto", NewCreateDto(contract))
	createDto := dtoAdapter.GetDtoAdapter("createDto")
	err := createDto.Convert()
	if err != nil {
		return nil, err
	}
	data := createDto.Getter()
	entity, ok := data.(entity.Contract)
	if !ok {
		return nil, GetError(http.StatusInternalServerError, errors.New("faile to assertion create DTO"))
	}
	return &entity, nil
}

// Concrete Update DTO
type updateDto struct {
	contract   *models.UpdateContract
	updateList map[string]any
}

func (updateDto *updateDto) doMap(key string, value any) {
	if value != nil {
		updateDto.updateList[key] = value
	}
}

func (updateDto *updateDto) mapStudentInfo() error {
	updateDto.doMap("ID", updateDto.contract.ID)
	updateDto.doMap("StudentCode", updateDto.contract.StudentCode)
	updateDto.doMap("FirstName", updateDto.contract.FirstName)
	updateDto.doMap("LastName", updateDto.contract.LastName)
	updateDto.doMap("MiddleName", updateDto.contract.MiddleName)
	updateDto.doMap("Email", updateDto.contract.Email)
	updateDto.doMap("Dob", updateDto.contract.Dob)
	updateDto.doMap("Gender", updateDto.contract.Gender)
	avatarString, err := decodeAvatar(updateDto.contract.Avatar)
	if err != nil {
		return GetError(http.StatusUnprocessableEntity, err)
	}
	updateDto.doMap("Avatar", avatarString)

	return nil
}

func (updateDto *updateDto) mapContractInfo() {
	updateDto.doMap("Sign", updateDto.contract.Sign)
	updateDto.doMap("Phone", updateDto.contract.Phone)
	updateDto.doMap("IsActive", updateDto.contract.IsActive)
	updateDto.doMap("Address", updateDto.contract.Address)
	updateDto.doMap("RoomID", updateDto.contract.RoomID)
	updateDto.doMap("NotificationChannels", updateDto.contract.NotificationChannels)
}

func (updateDto *updateDto) Convert() error {
	err := updateDto.mapStudentInfo()
	if err != nil {
		return err
	}
	updateDto.mapContractInfo()

	return nil
}

func (updateDto *updateDto) Getter() any {
	return updateDto
}

func NewUpdateDto(contract *models.UpdateContract) *updateDto {
	return &updateDto{contract: contract}
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

// Concrete Get and List DTO
type getListDto struct {
	contract *models.ReplyContract
	entity   *entity.Contract
}

func (getListDto *getListDto) convertStudentInfo() {
	getListDto.contract.Id = &getListDto.entity.ID
	getListDto.contract.StudentCode = &getListDto.entity.StudentCode
	getListDto.contract.FirstName = &getListDto.entity.FirstName
	getListDto.contract.LastName = &getListDto.entity.LastName
	getListDto.contract.MiddleName = &getListDto.entity.MiddleName
	getListDto.contract.Email = &getListDto.entity.Email
	gender := uint32(getListDto.entity.Gender)
	getListDto.contract.Gender = &gender
	avatar := encodeAvatar(getListDto.entity.Avatar)
	getListDto.contract.Avatar = &avatar
}

func (getListDto *getListDto) convertContractInfo() {
	getListDto.contract.Sign = &getListDto.entity.Sign
	getListDto.contract.Phone = &getListDto.entity.Phone
	getListDto.contract.Dob = &getListDto.entity.DOB
	getListDto.contract.Address = &getListDto.entity.Address
	getListDto.contract.IsActive = &getListDto.entity.IsActive
	getListDto.contract.RegistryAt = &getListDto.entity.RegistryAt
	getListDto.contract.RoomId = &getListDto.entity.RoomID
	notification := uint32(getListDto.entity.NotificationChannels)
	getListDto.contract.NotificationChannels = &notification
}

func (getListDto *getListDto) Convert() error {
	getListDto.convertStudentInfo()
	getListDto.convertContractInfo()
	return nil
}

func NewGetListDto(entity *entity.Contract) *getListDto {
	return &getListDto{entity: entity}
}

func (getListDto *getListDto) Getter() any {
	return getListDto
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
	contract, ok := data.(models.ReplyContract)
	if !ok {
		return nil, GetError(http.StatusInternalServerError, errors.New("faile to assertion get list DTO"))
	}

	return &contract, nil
}
