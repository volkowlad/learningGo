package main

import (
	"encoding/json"
	"errors"
)

type Create struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	PassCon  string `json:"pass_con"`
}

var (
	ErrEmail = errors.New("email is required")
	ErrPas1  = errors.New("pass is req")
	ErrPas2  = errors.New("passcon is req")
	ErrPas3  = errors.New("pass != passcon")
)

func valid(req Create) error {
	if req.Email == "" {
		return ErrEmail
	}

	if req.Password == "" {
		return ErrPas1
	}

	if req.PassCon == "" {
		return ErrPas2
	}

	if req.Password != req.PassCon {
		return ErrPas3
	}
	return nil
}

func Decode(request []byte) (Create, error) {
	req := Create{}

	err := json.Unmarshal(request, req)
	err = valid(req)
	if err != nil {
		return Create{}, err
	}

	return req, nil
}
