package userBiz

import (
	"GoCare/common"
	"GoCare/components/tokenprovider"
	userModel "GoCare/module/user/model"
	"context"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
}

type loginBusiness struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// Login flow: find user's email -> hash pass from input and compare with db
// then provider issue JWT token for client, return token

func (business *loginBusiness) Login(ctx context.Context, data *userModel.UserLogin) (*tokenprovider.Token, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, userModel.ErrEmailOrPasswordInvalid
	}

	passHashed := business.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, userModel.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayLoad{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrorInternal(err)
	}
	return accessToken, nil
}
