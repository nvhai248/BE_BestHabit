package userbiz

import (
	"bestHabit/common"
	"bestHabit/modules/user/usermodel"
	"context"
)

type DeleteUserStore interface {
	DeleteUser(ctx context.Context, userId int) error
	FindById(ctx context.Context, id int) (*usermodel.UserFind, error)
}

type deleteUserBiz struct {
	store DeleteUserStore
}

func NewDeleteUserBiz(store DeleteUserStore) *deleteUserBiz {
	return &deleteUserBiz{store: store}
}

func (b *deleteUserBiz) DeleteUser(ctx context.Context, userId int) error {
	user, err := b.store.FindById(ctx, userId)

	if err != nil {
		if err == common.ErrorNoRows {
			return common.ErrEntityNotFound(usermodel.EntityName, err)
		}
		return err
	}

	if user.Status == common.UserDeleted {
		return common.ErrEntityDeleted(usermodel.EntityName, err)
	}

	if err := b.store.DeleteUser(ctx, userId); err != nil {
		return err
	}

	return nil
}
