package biz

import (
	"context"
	"social-photo/common"
	"social-photo/component/tokenprovider"
	"social-photo/modules/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*model.User, error)
}

type loginBiz struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBiz {
	return &loginBiz{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *loginBiz) Login(ctx context.Context, data *model.UserLogin) (tokenprovider.Token, error) {

	// 1. Find user, email
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	// 2. Hash password from input and compare with pass in db
	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	// 3. Provider: issue JWT token for client
	payload := &common.TokenPayload{
		UId:   user.Id,
		URole: user.Role.String(),
	}

	// 3.1. Access token and refresh token
	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	// 4. Return token(s)
	return accessToken, nil
}
