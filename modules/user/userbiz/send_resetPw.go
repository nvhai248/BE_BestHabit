package userbiz

import (
	"bestHabit/common"
	"bestHabit/component/mailprovider"
	"bestHabit/component/tokenprovider"
	"bestHabit/modules/user/usermodel"
	"context"
	"fmt"
)

type SendResetPwStore interface {
	FindByEmail(ctx context.Context, email string) (*usermodel.UserFind, error)
}

type sendResetPwBiz struct {
	store         SendResetPwStore
	mailSender    mailprovider.EmailSender
	tokenProvider tokenprovider.Provider
}

func NewSendResetPwBiz(
	store SendResetPwStore,
	mailSender mailprovider.EmailSender,
	tokenProvider tokenprovider.Provider,
) *sendResetPwBiz {
	return &sendResetPwBiz{mailSender: mailSender, store: store, tokenProvider: tokenProvider}
}

func (biz *sendResetPwBiz) SendResetPw(ctx context.Context, email common.ToUser) error {
	user, err := biz.store.FindByEmail(ctx, email.Email)

	if err != nil {
		if err == common.ErrorNoRows {
			return common.ErrEntityNotFound(usermodel.EntityName, err)
		}

		return err
	}

	if user.Status == common.UserDeleted {
		return common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	go func() {
		defer common.AppRecover()
		payload := tokenprovider.TokenPayload{
			UserId: user.Id,
			Role:   *user.Role,
		}
		token, _err := biz.tokenProvider.Generate(payload, 60*60*24)

		if _err != nil {
			fmt.Println(_err)
		}

		email := common.NewRequireResetPw([]string{email.Email}, token.Token)
		biz.mailSender.SendEmail(email.Subject, email.Content, email.To, email.Cc, email.Bcc, email.AttachFiles)
	}()

	return nil
}
