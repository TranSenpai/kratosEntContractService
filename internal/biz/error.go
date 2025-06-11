package biz

import kerror "github.com/go-kratos/kratos/v2/errors"

func GetError(value int, err error) error {
	return kerror.New(value, "Biz error", err.Error())
}
