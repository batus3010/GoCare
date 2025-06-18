package userBiz

import (
	"GoCare/common"
	userModel "GoCare/module/user/model"
	"context"
)

// RegisterStorage requires finding if the user is already exist with email first
type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
	CreateUser(ctx context.Context, data *userModel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{registerStorage: registerStorage, hasher: hasher}
}

func (business *registerBusiness) Register(ctx context.Context, data *userModel.UserCreate) error {
	user, _ := business.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		return userModel.ErrEmailExisted
	}

	salt := common.GenSalt(50)
	data.Password = business.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "receptionist" // TODO: remove hard code

	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrorCannotCreateEntity(userModel.EntityUser, err)
	}

	return nil
}
