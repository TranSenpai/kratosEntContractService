package biz

import (
	"context"
	models "dormitory/internal/models"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (c contractBiz) checkStudentCode(studentCode string) string {
	if len(studentCode) != 10 {
		return "must input student code | "
	}

	return ""
}

func (c contractBiz) checkFirstLastName(name string) string {
	// In Go, a string is a squence of bytes.
	// loop for each runes of a string, rune is like char
	// var a = 'A' -> a this rune type
	if name == "" {
		return "please input first name and last name | "
	}
	for _, v := range name {
		if v >= 65 && v <= 90 || v >= 97 && v <= 122 {
			continue
		}
		return "first name and last name can not contain special character | "
	}

	return ""
}

func (c contractBiz) checkEmail(email string) string {
	if email == "" {
		return "must input email | "
	}
	// _, err := mail.ParseAddress(email)
	// if err != nil {
	// 	return GetError(http.StatusUnprocessableEntity, errors.New("invalid email"))
	// }

	return ""
}

func (c contractBiz) checkPhone(phone string) string {
	if len(phone) == 0 {
		return "must input phone | "
	}
	if len(phone) != 10 {
		return "invalid phone | "
	}

	return ""
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

func (c contractBiz) checkRoom(ctx context.Context, roomID string) string {
	if roomID == "" {
		return "must input room | "
	}
	if len(roomID) != 5 {
		return "invalid input room, room must be 5 character, first character is building, next two character is floor, last two character is room | "
	}
	for k, v := range roomID {
		check := isValidRoom(k, v)
		if !check {
			return "invalid input room, room must be 5 character, first character is building, next two character is floor, last two character is room | "
		}
	}
	totalContract, err := c.contractRepo.GetTotalContractRoom(ctx, roomID)
	if err != nil {
		return err.Error()
	}
	if totalContract.Total >= 4 {
		return "room full | "
	}

	return ""
}

func (c contractBiz) checkRegisterTime(registerTime time.Time) string {
	checkRegisterTime := time.Date(registerTime.Year(), registerTime.Month(), registerTime.Day(), 19, 0, 0, 0, time.UTC)
	fmt.Println(checkRegisterTime)
	if registerTime.After(checkRegisterTime) {
		return "don't create contract after 7pm please | "
	}

	return ""
}

func (c contractBiz) checkRequiredField(ctx context.Context, contractModel *models.CreateContract) error {
	var strError string
	if contractModel == nil {
		return GetError(http.StatusBadRequest, errors.New("nil contract model"))
	}
	strError += c.checkStudentCode(contractModel.StudentCode)
	strError += c.checkFirstLastName(contractModel.FirstName)
	strError += c.checkFirstLastName(contractModel.LastName)
	strError += c.checkEmail(contractModel.Email)
	strError += c.checkPhone(contractModel.Phone)
	strError += c.checkRoom(ctx, contractModel.RoomId)
	strError += c.checkRegisterTime(contractModel.RegistryAt)
	if strError == "" {
		return nil
	}
	strError = strError[:len(strError)-3]

	return GetError(http.StatusBadRequest, errors.New(strError))
}
