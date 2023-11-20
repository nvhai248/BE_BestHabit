package userbiz

import (
	"bestHabit/common"
	"bestHabit/modules/user/usermodel"
	"context"
)

type BasicRegisterStorage interface {
	FindByEmail(ctx context.Context, email string) (*usermodel.UserFind, error)
	Create(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type basicRegisterBiz struct {
	store  BasicRegisterStorage
	hasher Hasher
}

func NewBasicRegisterBiz(store BasicRegisterStorage, hasher Hasher) *basicRegisterBiz {
	return &basicRegisterBiz{store: store, hasher: hasher}
}

func (biz *basicRegisterBiz) BasicRegister(ctx context.Context, data *usermodel.UserCreate) error {

	user, err := biz.store.FindByEmail(ctx, data.Email)

	if err != common.ErrorNoRows && err != nil {
		return err
	}

	if user != nil {
		return common.ErrEmailExisted
	}

	if err := data.Validate(); err != nil {
		return usermodel.ErrNameCannotBeEmpty
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)

	data.Salt = salt
	data.Settings = common.NewDefaultSettings()
	data.Role = "user" // hard code

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
