package userbiz

import (
	"bestHabit/common"
	"bestHabit/component/mailprovider"
	"bestHabit/component/tokenprovider"
	"context"
	"fmt"
)

type sendVerificationBiz struct {
	mailSender    mailprovider.EmailSender
	tokenProvider tokenprovider.Provider
}

func NewSendVerificationBiz(
	mailSender mailprovider.EmailSender,
	tokenProvider tokenprovider.Provider,
) *sendVerificationBiz {
	return &sendVerificationBiz{mailSender: mailSender, tokenProvider: tokenProvider}
}

func (biz *sendVerificationBiz) SendVerification(ctx context.Context, email string, id int, role string) error {

	go func() {
		defer common.AppRecover()
		payload := tokenprovider.TokenPayload{
			UserId: id,
			Role:   role,
		}
		token, _err := biz.tokenProvider.Generate(payload, 60*60*24)

		if _err != nil {
			fmt.Println(_err)
		}

		email := common.NewEmailVerifyAccount([]string{email}, token.Token)
		biz.mailSender.SendEmail(email.Subject, email.Content, email.To, email.Cc, email.Bcc, email.AttachFiles)
	}()

	return nil
}
