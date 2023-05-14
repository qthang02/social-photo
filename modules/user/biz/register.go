package biz

import (
	"context"
	"social-photo/common"
	"social-photo/modules/user/model"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*model.User, error)
	CreateUser(ctx context.Context, data *model.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBiz(storage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{registerStorage: storage, hasher: hasher}
}

func (biz *registerBusiness) Register(ctx context.Context, data *model.UserCreate) error {
	// check user is exited
	user, _ := biz.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return model.ErrEmailExisted
	}

	// create user

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	data.Status = 1

	if err := biz.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
