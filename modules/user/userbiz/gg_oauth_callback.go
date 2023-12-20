package userbiz

import (
	"bestHabit/common"
	"bestHabit/component/mailprovider"
	"bestHabit/component/tokenprovider"
	"bestHabit/modules/user/usermodel"
	"context"
	"fmt"
)

type GgOauthCallbackStorage interface {
	FindByEmail(ctx context.Context, email string) (*usermodel.UserFind, error)
	FindByGgId(ctx context.Context, ggId string) (*usermodel.UserFind, error)
	Create(ctx context.Context, data *usermodel.UserCreate) error
}

type ggOauthCallbackBiz struct {
	store         GgOauthCallbackStorage
	mailSender    mailprovider.EmailSender
	tokenProvider tokenprovider.Provider
}

func NewGgOauthCallbackBiz(
	store GgOauthCallbackStorage,
	mailSender mailprovider.EmailSender,
	tokenProvider tokenprovider.Provider,
) *ggOauthCallbackBiz {
	return &ggOauthCallbackBiz{store: store, mailSender: mailSender, tokenProvider: tokenProvider}
}

func (b *ggOauthCallbackBiz) GgOauthCallback(ctx context.Context, user *usermodel.UserCreate) (*tokenprovider.Token, error) {
	// find user by gg id
	userCheck, err := b.store.FindByGgId(ctx, user.GgID)

	if err == common.ErrorNoRows {
		// if user not found => find user by email
		userCheckByMail, err := b.store.FindByEmail(ctx, user.Email)

		if err == common.ErrorNoRows {
			// if not found => create new user and return
			err := b.store.Create(ctx, user)

			if err != nil {
				return nil, common.NewCustomError(err,
					"Cannot create user!",
					"CannotCreateUser")
			}

			userReturn, err := b.store.FindByGgId(ctx, user.GgID)

			if err != nil {
				return nil, common.NewCustomError(err,
					"Cannot find user after create user!",
					"CannotFindUser")
			}

			payload := tokenprovider.TokenPayload{
				UserId: userReturn.Id,
				Role:   *userReturn.Role,
			}

			accessToken, err := b.tokenProvider.Generate(payload, 60*60*24*7)
			if err != nil {
				return nil, common.ErrInternal(err)
			}

			go func() {
				defer common.AppRecover()
				token, err := b.tokenProvider.Generate(payload, 60*60*24)
				if err != nil {
					fmt.Print("Error generating token!")
				}
				email := common.NewEmailVerifyAccount([]string{*userReturn.Email}, token.Token)
				b.mailSender.SendEmail(email.Subject, email.Content, email.To, email.Cc, email.Bcc, email.AttachFiles)
			}()

			return accessToken, nil
		}

		if err != nil {
			return nil, common.ErrInternal(err)
		}

		// if found => return err
		return nil, common.NewCustomError(err,
			fmt.Sprintf("The email %s already used by another account!", *userCheckByMail.Email),
			"CannotFindUser")
	}

	// if find => return token
	payload := tokenprovider.TokenPayload{
		UserId: userCheck.Id,
		Role:   *userCheck.Role,
	}

	accessToken, err := b.tokenProvider.Generate(payload, 60*60*24*7)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
