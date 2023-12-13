package userbiz

import (
	"bestHabit/common"
	"bestHabit/modules/user/usermodel"
	"context"
)

type ResetPasswordStore interface {
	FindById(ctx context.Context, id int) (*usermodel.UserFind, error)
	ChangePassword(ctx context.Context, newPw string, userId int) error
}

type resetPasswordBiz struct {
	store  ResetPasswordStore
	hasher Hasher
}

func NewResetPasswordBiz(
	store ResetPasswordStore,
	hasher Hasher,
) *resetPasswordBiz {
	return &resetPasswordBiz{store: store, hasher: hasher}
}

func (biz *resetPasswordBiz) ResetPassword(ctx context.Context, newPw *usermodel.ResetPassword, id int) error {
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

	newPass := biz.hasher.Hash(*newPw.Password + *user.Salt)

	if err := biz.store.ChangePassword(ctx, newPass, user.Id); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, nil)
	}

	return nil
}
