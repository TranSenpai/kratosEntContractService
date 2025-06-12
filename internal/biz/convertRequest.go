package biz

import (
	entity "dormitory/internal/entities"
	models "dormitory/internal/models"
	"encoding/base64"
	"errors"
	"net/http"
)

func (c contractBiz) convertCreateStudentInfo(entity *entity.Contract, contract *models.CreateContract) error {
	if entity == nil || contract == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("create contract or entity nil"))
	}
	entity.StudentCode = contract.StudentCode
	entity.FirstName = contract.FirstName
	entity.LastName = contract.LastName
	entity.MiddleName = contract.MiddleName
	entity.Email = contract.Email
	entity.Gender = uint8(contract.Gender)
	if contract.Dob != nil {
		entity.DOB = *contract.Dob
	}
	avatarByte, err := decodeAvatar(&contract.Avatar)
	if err != nil {
		return GetError(http.StatusUnprocessableEntity, err)
	}
	entity.Avatar = avatarByte

	return nil
}

func (c contractBiz) convertCreateContractInfo(entity *entity.Contract, contract *models.CreateContract) error {
	if entity == nil || contract == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("create contract or entity nil"))
	}
	entity.Sign = contract.Sign
	entity.Phone = contract.Phone
	entity.Address = contract.Address
	entity.IsActive = contract.IsActive
	entity.RoomID = contract.RoomId
	entity.NotificationChannels = uint8(contract.NotificationChannels)

	return nil
}

func (c contractBiz) convertCreateContractRequest(contract *models.CreateContract) (*entity.Contract, error) {
	entity := &entity.Contract{}
	if contract == nil {
		return nil, GetError(http.StatusUnprocessableEntity, errors.New("create contract or entity nil"))
	}
	err := c.convertCreateStudentInfo(entity, contract)
	if err != nil {
		return nil, err
	}
	err = c.convertCreateContractInfo(entity, contract)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (c contractBiz) doMap(updateList map[string]any, key string, value any) {
	if value != nil {
		updateList[key] = value
	}
}

func (c contractBiz) mapStudentInfo(updateList map[string]any, contract *models.UpdateContract) error {
	if contract == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("update contract nil"))
	}
	c.doMap(updateList, "ID", contract.ID)
	c.doMap(updateList, "StudentCode", contract.StudentCode)
	c.doMap(updateList, "FirstName", contract.FirstName)
	c.doMap(updateList, "LastName", contract.LastName)
	c.doMap(updateList, "MiddleName", contract.MiddleName)
	c.doMap(updateList, "Email", contract.Email)
	c.doMap(updateList, "Dob", contract.Dob)
	c.doMap(updateList, "Gender", contract.Gender)
	avatarString, err := decodeAvatar(contract.Avatar)
	if err != nil {
		return GetError(http.StatusUnprocessableEntity, err)
	}
	c.doMap(updateList, "Avatar", avatarString)

	return nil
}

func (c contractBiz) mapContractInfo(updateList map[string]any, contract *models.UpdateContract) error {
	if contract == nil || updateList == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("update contract nil"))
	}
	c.doMap(updateList, "Sign", contract.Sign)
	c.doMap(updateList, "Phone", contract.Phone)
	c.doMap(updateList, "IsActive", contract.IsActive)
	c.doMap(updateList, "Address", contract.Address)
	c.doMap(updateList, "RoomID", contract.RoomID)
	c.doMap(updateList, "NotificationChannels", contract.NotificationChannels)

	return nil
}

func (c contractBiz) convertUpdateContractRequest(contract *models.UpdateContract) (map[string]any, error) {
	updateList := make(map[string]any)
	if contract == nil {
		return nil, GetError(http.StatusUnprocessableEntity, errors.New("update contract nil"))
	}
	err := c.mapStudentInfo(updateList, contract)
	if err != nil {
		return nil, err
	}
	err = c.mapContractInfo(updateList, contract)
	if err != nil {
		return nil, err
	}

	return updateList, nil
}

func (c contractBiz) convertReplyStudentInfo(replyContract *models.ReplyContract, contract *entity.Contract) error {
	if replyContract == nil || contract == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("reply contract or entity nil"))
	}
	replyContract.Id = &contract.ID
	replyContract.StudentCode = &contract.StudentCode
	replyContract.FirstName = &contract.FirstName
	replyContract.LastName = &contract.LastName
	replyContract.MiddleName = &contract.MiddleName
	replyContract.Email = &contract.Email
	gender := uint32(contract.Gender)
	replyContract.Gender = &gender
	avatar := encodeAvatar(contract.Avatar)
	replyContract.Avatar = &avatar

	return nil
}

func (c contractBiz) convertReplyContractInfo(replyContract *models.ReplyContract, contract *entity.Contract) error {
	if replyContract == nil || contract == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("reply contract or entity nil"))
	}
	replyContract.Sign = &contract.Sign
	replyContract.Phone = &contract.Phone
	replyContract.Dob = &contract.DOB
	replyContract.Address = &contract.Address
	replyContract.IsActive = &contract.IsActive
	replyContract.RegistryAt = &contract.RegistryAt
	replyContract.RoomId = &contract.RoomID
	notification := uint32(contract.NotificationChannels)
	replyContract.NotificationChannels = &notification

	return nil
}

func (c contractBiz) convertReplyContract(contract *entity.Contract) (*models.ReplyContract, error) {
	replyContract := &models.ReplyContract{}
	err := c.convertReplyStudentInfo(replyContract, contract)
	if err != nil {
		return nil, err
	}
	err = c.convertReplyContractInfo(replyContract, contract)
	if err != nil {
		return nil, err
	}
	return replyContract, nil
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
