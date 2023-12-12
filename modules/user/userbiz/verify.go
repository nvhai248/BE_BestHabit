package userbiz

import (
	"bestHabit/common"
	"bestHabit/modules/user/usermodel"
	"context"
)

type VerifyUserStore interface {
	FindById(ctx context.Context, id int) (*usermodel.UserFind, error)
	VerifyUser(ctx context.Context, userId int) error
}

type verifyUserBiz struct {
	store VerifyUserStore
}

func NewVerifyUserBiz(store VerifyUserStore) *verifyUserBiz {
	return &verifyUserBiz{store: store}
}

func (biz *verifyUserBiz) VerifyBiz(ctx context.Context, id int) error {
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

	if err := biz.store.VerifyUser(ctx, id); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}
