package biz

import (
	entity "dormitory/internal/ent"
	models "dormitory/internal/models"
	"encoding/base64"
	"errors"
	"net/http"
	"time"
)

func (c contractBiz) convertCreateStudentInfo(entity *entity.Contract, contract *models.CreateContract) error {
	if entity == nil || contract == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("create contract or ent nil"))
	}
	entity.StudentCode = contract.StudentCode
	entity.FirstName = contract.FirstName
	entity.LastName = contract.LastName
	entity.MiddleName = contract.MiddleName
	entity.Email = contract.Email
	entity.Gender = uint8(contract.Gender)
	if contract.Dob != nil {
		entity.Dob = *contract.Dob
	}
	avatarByte, err := c.decodeAvatar(&contract.Avatar)
	if err != nil {
		return GetError(http.StatusUnprocessableEntity, err)
	}
	entity.Avatar = avatarByte

	return nil
}

func (c contractBiz) convertCreateContractInfo(entity *entity.Contract, contract *models.CreateContract) error {
	if entity == nil || contract == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("create contract or ent nil"))
	}
	entity.Sign = contract.Sign
	entity.Phone = contract.Phone
	entity.Address = contract.Address
	entity.IsActive = contract.IsActive
	entity.RoomID = contract.RoomId
	entity.NotificationChannels = uint8(contract.NotificationChannels)
	entity.RegistryAt = contract.RegistryAt

	return nil
}

func (c contractBiz) convertCreateContractRequest(contract *models.CreateContract) (*entity.Contract, error) {
	if contract == nil {
		return nil, GetError(http.StatusUnprocessableEntity, errors.New("create contract or ent nil"))
	}
	entity := &entity.Contract{}
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

func (c contractBiz) mapStudentInfo(updateList map[string]any, contract *models.UpdateContract) error {
	if contract == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("update contract nil"))
	}
	if contract.ID != nil {
		updateList["ID"] = *contract.ID
	}
	if contract.StudentCode != nil {
		updateList["StudentCode"] = *contract.StudentCode
	}
	if contract.FirstName != nil {
		updateList["FirstName"] = *contract.FirstName
	}
	if contract.LastName != nil {
		updateList["LastName"] = *contract.LastName
	}
	if contract.MiddleName != nil {
		updateList["MiddleName"] = *contract.MiddleName
	}
	if contract.Email != nil {
		updateList["Email"] = *contract.Email
	}
	if contract.Dob != nil {
		if time.Time.IsZero(*contract.Dob) {
			updateList["Dob"] = *contract.Dob
		}
	}
	if contract.Gender != nil {
		gender := *contract.Gender
		updateList["Gender"] = uint8(gender)
	}
	avatarByte, err := c.decodeAvatar(contract.Avatar)
	if err != nil {
		return GetError(http.StatusUnprocessableEntity, err)
	}
	if contract.Avatar != nil {
		updateList["Avatar"] = avatarByte
	}

	return nil
}

func (c contractBiz) mapContractInfo(updateList map[string]any, contract *models.UpdateContract) error {
	if contract == nil || updateList == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("update contract nil"))
	}
	if contract.Sign != nil {
		updateList["Sign"] = *contract.Sign
	}
	if contract.Phone != nil {
		updateList["Phone"] = *contract.Phone
	}
	if contract.IsActive != nil {
		updateList["IsActive"] = *contract.IsActive
	}
	if contract.Address != nil {
		updateList["Address"] = *contract.Address
	}
	if contract.RoomID != nil {
		updateList["RoomID"] = *contract.RoomID
	}
	if contract.NotificationChannels != nil {
		updateList["NotificationChannels"] = *contract.NotificationChannels
	}

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
		return GetError(http.StatusUnprocessableEntity, errors.New("reply contract or ent nil"))
	}
	replyContract.Id = contract.ID
	replyContract.StudentCode = contract.StudentCode
	replyContract.FirstName = contract.FirstName
	replyContract.LastName = contract.LastName
	replyContract.MiddleName = contract.MiddleName
	replyContract.Email = contract.Email
	gender := uint32(contract.Gender)
	replyContract.Gender = gender
	avatarStr := c.encodeAvatar(contract.Avatar)
	replyContract.Avatar = avatarStr

	return nil
}

func (c contractBiz) convertReplyContractInfo(replyContract *models.ReplyContract, contract *entity.Contract) error {
	if replyContract == nil || contract == nil {
		return GetError(http.StatusUnprocessableEntity, errors.New("reply contract or ent nil"))
	}
	replyContract.Sign = contract.Sign
	replyContract.Phone = contract.Phone
	replyContract.Dob = contract.Dob
	replyContract.Address = contract.Address
	replyContract.IsActive = contract.IsActive
	replyContract.RegistryAt = contract.RegistryAt
	replyContract.RoomId = contract.RoomID
	notification := uint32(contract.NotificationChannels)
	replyContract.NotificationChannels = notification

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

func (c contractBiz) decodeAvatar(avatar *string) ([]byte, error) {
	if avatar != nil && *avatar != "" {
		avatarDecoded, err := base64.StdEncoding.DecodeString(*avatar)
		if err != nil {
			return nil, GetError(http.StatusUnprocessableEntity, err)
		}
		return avatarDecoded, nil
	}

	return nil, nil
}

func (c contractBiz) encodeAvatar(avatar []byte) string {
	if avatar != nil {
		avatarDecoded := base64.StdEncoding.EncodeToString(avatar)
		return avatarDecoded
	}

	return ""
}
