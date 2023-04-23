package tokenprovider

import (
	"Food-Delivery/common"
	"errors"
	"time"
)

type Provider interface {
	Generate(data TokenPayLoad, expiry int) (*Token, error) // call func trong folder jwt
	Validate(token string) (*TokenPayLoad, error)
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)

	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)
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
