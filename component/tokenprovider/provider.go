package tokenprovider

import (
	"errors"
	"social-photo/common"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (Token, error)
	Validate(token string) (TokenPayload, error)
	SecretKey() string
}

type TokenPayload interface {
	UserId() int
	Role() string
}

type Token interface {
	GetToken() string
}

var (
	ErrNotFound      = common.NewCustomError(errors.New("token not found"), "token not found", "ErrNotFound")
	ErrEncodingToken = common.NewCustomError(errors.New("error encoding token"), "error encoding token", "ErrEncodingToken")
	ErrInvalidToken  = common.NewCustomError(errors.New("error invalid token"), "error invalid token", "ErrInvalidToken")
)
