package main

import (
	"errors"
)

// некретичная ошибка валидации
type NonCrit struct{}

func (e NonCrit) Error() string {
	return "valid error"
}

// критичные ошибки
var (
	ErrCon = errors.New("con")
	ErrReq = errors.New("req")
)

const Unknow = "unknow error"

func GetMsg(err error) string {
	if errors.Is(err, ErrCon) {

		return err.Error()

	} else if errors.Is(err, ErrReq) {

		return err.Error()

	} else if errors.As(err, &NonCrit{}) {

		return ""

	}

	return Unknow
}
