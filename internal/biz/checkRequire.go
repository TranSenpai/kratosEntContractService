package biz

import (
	"context"
	models "dormitory/internal/models"
	"errors"
	"net/http"
)

func (c contractBiz) checkStudentCode(studentCode string) error {
	if len(studentCode) != 10 {
		return GetError(http.StatusUnprocessableEntity, errors.New("must input student code"))
	}

	return nil
}

func (c contractBiz) checkFirstName(firstName string) error {
	// In Go, a string is a squence of bytes.
	// loop for each runes of a string, rune is like char
	// var a = 'A' -> a this rune type
	for _, v := range firstName {
		if v < 65 && v > 90 || v < 61 && v > 122 {
			return GetError(http.StatusUnprocessableEntity, errors.New("first name can not contain special character"))
		}
	}

	return nil
}

func (c contractBiz) checkLastName(lastName string) error {
	for _, v := range lastName {
		if v < 65 && v > 90 || v < 61 && v > 122 {
			return GetError(http.StatusUnprocessableEntity, errors.New("last name can not contain special character"))
		}
	}

	return nil
}

func (c contractBiz) checkEmail(email string) error {
	if email == "" {
		return GetError(http.StatusUnprocessableEntity, errors.New("must input email"))
	}
	// _, err := mail.ParseAddress(email)
	// if err != nil {
	// 	return GetError(http.StatusUnprocessableEntity, errors.New("invalid email"))
	// }

	return nil
}

func (c contractBiz) checkPhone(phone string) error {
	if len(phone) != 10 {
		return GetError(http.StatusUnprocessableEntity, errors.New("invalid phone"))
	}

	return nil
}

func isValidRoom(k int, v rune) bool {
	flag := true
	if k == 0 && v < 65 && v > 90 || v < 61 && v > 122 {
		flag = false
	} else if k != 0 && v < 48 && v > 57 {
		flag = false
	}

	return flag
}

func (c contractBiz) checkRoom(ctx context.Context, roomID string) error {
	if roomID == "" {
		return GetError(http.StatusUnprocessableEntity, errors.New("must input room"))
	}
	if len(roomID) != 5 {
		return GetError(http.StatusBadRequest, errors.New("invalid input room, room must be 5 character, first character is building, next two character is floor, last two character is room"))
	}
	for k, v := range roomID {
		check := isValidRoom(k, v)
		if !check {
			return GetError(http.StatusBadRequest, errors.New("invalid input room, room must be 5 character, first character is building, next two character is floor, last two character is room"))
		}
	}
	totalContract, err := c.contractRepo.GetTotalContractRoom(ctx, roomID)
	if err != nil {
		return err
	}
	if totalContract.Total > 4 {
		return GetError(http.StatusBadRequest, errors.New("room full"))
	}

	return nil
}

func (c contractBiz) CheckRequiredField(ctx context.Context, contractModel *models.CreateContract) error {
	if err := c.checkStudentCode(contractModel.StudentCode); err != nil {
		return err
	}
	if err := c.checkFirstName(contractModel.FirstName); err != nil {
		return err
	}
	if err := c.checkLastName(contractModel.LastName); err != nil {
		return err
	}
	if err := c.checkEmail(contractModel.Email); err != nil {
		return err
	}
	if err := c.checkPhone(contractModel.Phone); err != nil {
		return err
	}
	if err := c.checkRoom(ctx, contractModel.RoomId); err != nil {
		return err
	}

	return nil
}
