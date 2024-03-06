package userbiz

import (
	"bestHabit/common"
	"bestHabit/component/mailprovider"
	"bestHabit/component/tokenprovider"
	"bestHabit/modules/user/usermodel"
	"context"
	"fmt"
)

type ChangePassStore interface {
	FindById(ctx context.Context, id int) (*usermodel.UserFind, error)
	ChangePassword(ctx context.Context, newPw string, userId int) error
}

type ChangePassGRPCStore interface {
	UpdatePasswordByGRPC(ctx context.Context, userId int, newPassword string) (*int, error)
}

type changePassBiz struct {
	store         ChangePassStore
	hasher        Hasher
	mailSender    mailprovider.EmailSender
	tokenProvider tokenprovider.Provider
	gRPCStore     ChangePassGRPCStore
}

func NewChangePassBiz(
	store ChangePassStore,
	hasher Hasher,
	mailSender mailprovider.EmailSender,
	tokenProvider tokenprovider.Provider,
	gRPCStore ChangePassGRPCStore,
) *changePassBiz {
	return &changePassBiz{mailSender: mailSender,
		hasher: hasher, store: store,
		tokenProvider: tokenProvider,
		gRPCStore:     gRPCStore}
}

func (biz *changePassBiz) ChangePass(ctx context.Context, email string, id int, role string, newPw *usermodel.UpdatePassword) error {
	user, err := biz.store.FindById(ctx, id)

	if err != nil {
		if err == common.ErrorNoRows {
			return common.ErrEntityNotFound(usermodel.EntityName, err)
		}

		return err
	}

	if user.Status == common.UserDeleted {
		return common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	passHashed := biz.hasher.Hash(*newPw.Password + *user.Salt)

	if *user.Password != passHashed {
		return common.ErrEmailOrPasswordInvalid
	}

	newPass := biz.hasher.Hash(*newPw.NewPassword + *user.Salt)

	/* if err := biz.store.ChangePassword(ctx, newPass, id); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, nil)
	} */

	_, err = biz.gRPCStore.UpdatePasswordByGRPC(ctx, id, newPass)

	if err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

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

		email := common.NewRequireResetPwAfterChangePass([]string{email}, token.Token)
		biz.mailSender.SendEmail(email.Subject, email.Content, email.To, email.Cc, email.Bcc, email.AttachFiles)
	}()

	return nil
}
