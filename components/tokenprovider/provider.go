package tokenprovider

import (
	"GoCare/common"
	"errors"
	"time"
)

type Provider interface {
	Generate(data TokenPayLoad, expiry int) (*Token, error)
	Validate(token string) (*TokenPayLoad, error)
}

var (
	ErrNotFound = common.NewCustomErrorResponse(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound")

	ErrEncodingToken = common.NewCustomErrorResponse(
		errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken")

	ErrInvalidToken = common.NewCustomErrorResponse(
		errors.New("invalid token provided"),
		"ErrInvalidToken",
		"ErrInvalidToken")
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayLoad struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}
