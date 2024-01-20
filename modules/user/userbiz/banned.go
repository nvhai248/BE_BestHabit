package userbiz

import (
	"bestHabit/common"
	"bestHabit/modules/user/usermodel"
	"context"
)

type BannedUserStore interface {
	BannedUser(ctx context.Context, userId int) error
	FindById(ctx context.Context, id int) (*usermodel.UserFind, error)
}

type bannedUserBiz struct {
	store BannedUserStore
}

func NewBannedUserBiz(store BannedUserStore) *bannedUserBiz {
	return &bannedUserBiz{store: store}
}

func (b *bannedUserBiz) BannedUser(ctx context.Context, userId int) error {
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

	if user.Status == common.UserBanned {
		return common.NewCustomError(err, "User already banned!", "CannotBanned")
	}

	if err := b.store.BannedUser(ctx, userId); err != nil {
		return err
	}

	return nil
}
