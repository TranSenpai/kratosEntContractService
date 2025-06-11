package biz

import (
	"context"
	models "dormitory/internal/models"
	"errors"
	"net/http"
)

type ICheck interface {
	Result() error
}

type checks struct {
	kind map[string]ICheck
}

func (c checks) Setter(name string, kind ICheck) {
	c.kind[name] = kind
}

func (c checks) Getter(name string) ICheck {
	return c.kind[name]
}

func NewChecks(kind ICheck) *checks {
	return &checks{map[string]ICheck{}}
}

type checkRequiredField struct {
	err error
}

func (c checkRequiredField) checkStudentCode(studentCode string) error {
	if len(studentCode) != 10 {
		return GetError(http.StatusUnprocessableEntity, errors.New("must input student code"))
	}

	return nil
}

func (c checkRequiredField) checkFirstName(firstName string) error {
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

func (c checkRequiredField) checkLastName(lastName string) error {
	for _, v := range lastName {
		if v < 65 && v > 90 || v < 61 && v > 122 {
			return GetError(http.StatusUnprocessableEntity, errors.New("last name can not contain special character"))
		}
	}

	return nil
}

func (c checkRequiredField) checkEmail(email string) error {
	if email == "" {
		return GetError(http.StatusUnprocessableEntity, errors.New("must input email"))
	}
	// _, err := mail.ParseAddress(email)
	// if err != nil {
	// 	return GetError(http.StatusUnprocessableEntity, errors.New("invalid email"))
	// }

	return nil
}

func (c checkRequiredField) checkPhone(phone string) error {
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

func (c checkRequiredField) checkRoom(ctx context.Context, roomID string) error {
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
	var contractRepo IContractRepo
	totalContract, err := contractRepo.GetTotalContractRoom(ctx, roomID)
	if err != nil {
		return err
	}
	if totalContract.Total > 4 {
		return GetError(http.StatusBadRequest, errors.New("room full"))
	}

	return nil
}

func (c checkRequiredField) DoCheck(ctx context.Context, contractModel *models.CreateContract) {
	if err := c.checkStudentCode(contractModel.StudentCode); err != nil {
		c.err = errors.New(err.Error())
	}
	if err := c.checkFirstName(contractModel.FirstName); err != nil {
		c.err = errors.New(err.Error())
	}
	if err := c.checkLastName(contractModel.LastName); err != nil {
		c.err = errors.New(err.Error())
	}
	if err := c.checkEmail(contractModel.Email); err != nil {
		c.err = errors.New(err.Error())
	}
	if err := c.checkPhone(contractModel.Phone); err != nil {
		c.err = errors.New(err.Error())
	}
	if err := c.checkRoom(ctx, contractModel.RoomId); err != nil {
		c.err = errors.New(err.Error())
	}
}

func (c checkRequiredField) Result() error {
	return c.err
}

func NewCheckRequiredField() *checkRequiredField {
	return &checkRequiredField{}
}

func (c contractBiz) GetCheckRequiredField() ICheck {
	checkRequiredField := NewCheckRequiredField()
	checks := NewChecks(checkRequiredField)
	checks.Setter("checkRequiredField", checkRequiredField)
	return checks.Getter("checkRequiredField")
}
