package userbiz

import (
	"bestHabit/common"
	"bestHabit/modules/user/usermodel"
	"context"
)

type UnbannedUserStore interface {
	UnbannedUser(ctx context.Context, userId int) error
	FindById(ctx context.Context, id int) (*usermodel.UserFind, error)
}

type unbannedUserBiz struct {
	store UnbannedUserStore
}

func NewUnbannedUserBiz(store UnbannedUserStore) *unbannedUserBiz {
	return &unbannedUserBiz{store: store}
}

func (b *unbannedUserBiz) UnbannedUser(ctx context.Context, userId int) error {
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

	if user.Status == 1 {
		return common.NewCustomError(err, "User not banned!", "CannotUnbanned")
	}

	if err := b.store.UnbannedUser(ctx, userId); err != nil {
		return err
	}

	return nil
}
