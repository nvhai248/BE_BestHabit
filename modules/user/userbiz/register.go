package userbiz

import (
	"bestHabit/common"
	"bestHabit/component/mailprovider"
	"bestHabit/component/tokenprovider"
	"bestHabit/modules/user/usermodel"
	"context"
	"fmt"
)

type BasicRegisterStorage interface {
	FindByEmail(ctx context.Context, email string) (*usermodel.UserFind, error)
	Create(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type basicRegisterBiz struct {
	store         BasicRegisterStorage
	hasher        Hasher
	mailSender    mailprovider.EmailSender
	tokenProvider tokenprovider.Provider
}

func NewBasicRegisterBiz(
	store BasicRegisterStorage,
	hasher Hasher,
	mailSender mailprovider.EmailSender,
	tokenProvider tokenprovider.Provider,
) *basicRegisterBiz {
	return &basicRegisterBiz{store: store, hasher: hasher, mailSender: mailSender, tokenProvider: tokenProvider}
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
	data.Status = common.UserNotVerified

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	go func() {
		defer common.AppRecover()
		user, _err := biz.store.FindByEmail(ctx, data.Email)

		if _err != nil {
			fmt.Println("error: ", _err)
		}

		payload := tokenprovider.TokenPayload{
			UserId: user.Id,
			Role:   *user.Role,
		}
		token, _err := biz.tokenProvider.Generate(payload, 60*60*24)

		if _err != nil {
			return
		}

		email := common.NewEmailVerifyAccount([]string{data.Email}, token.Token)
		biz.mailSender.SendEmail(email.Subject, email.Content, email.To, email.Cc, email.Bcc, email.AttachFiles)
	}()

	return nil
}
