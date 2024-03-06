package userbiz

import (
	"bestHabit/common"
	"bestHabit/component/tokenprovider"
	"bestHabit/modules/user/usermodel"
	"context"
)

type BasicLoginStorage interface {
	FindByEmail(ctx context.Context, email string) (*usermodel.UserFind, error)
}

type TokenConfig interface {
	GetAtExp() int
}

type basicLoginBiz struct {
	storeUser     BasicLoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewBasicLoginBiz(
	storeUser BasicLoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *basicLoginBiz {
	return &basicLoginBiz{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// Process login
//1. Find Username, password
//2. Hash pass from input and compare with pass in db
//3. Provider: issue JWT token for client
//3.1. Access token and refresh token
//4. Return token(s)

func (biz *basicLoginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.storeUser.FindByEmail(ctx, data.Email)

	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	passHashed := biz.hasher.Hash(data.Password + *user.Salt)

	if *user.Password != passHashed {
		return nil, common.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   *user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
