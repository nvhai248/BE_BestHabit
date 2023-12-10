package userbiz

import (
	"bestHabit/common"
	"bestHabit/component/mailprovider"
	"bestHabit/component/tokenprovider"
	"bestHabit/modules/user/usermodel"
	"context"
	"fmt"
)

type SendVerificationStore interface {
	FindById(ctx context.Context, id int) (*usermodel.UserFind, error)
}

type sendVerificationBiz struct {
	store         SendVerificationStore
	mailSender    mailprovider.EmailSender
	tokenProvider tokenprovider.Provider
}

func NewSendVerificationBiz(
	store SendVerificationStore,
	mailSender mailprovider.EmailSender,
	tokenProvider tokenprovider.Provider,
) *sendVerificationBiz {
	return &sendVerificationBiz{store: store, mailSender: mailSender, tokenProvider: tokenProvider}
}

func (biz *sendVerificationBiz) SendVerification(ctx context.Context, id int, role string) error {

	go func() {
		defer common.AppRecover()
		user, _err := biz.store.FindById(ctx, id)

		if _err != nil {
			fmt.Println("error: ", _err)
		}

		payload := tokenprovider.TokenPayload{
			UserId: id,
			Role:   role,
		}
		token, _err := biz.tokenProvider.Generate(payload, 60*60*24)

		if _err != nil {
			fmt.Println(_err)
		}

		email := common.NewEmailVerifyAccount([]string{*user.Email}, token.Token)
		biz.mailSender.SendEmail(email.Subject, email.Content, email.To, email.Cc, email.Bcc, email.AttachFiles)
	}()

	return nil
}
