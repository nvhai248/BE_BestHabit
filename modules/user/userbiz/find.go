package userbiz

import (
	"bestHabit/common"
	"bestHabit/modules/user/usermodel"
	"context"
)

type FindUserStore interface {
	FindById(ctx context.Context, id int) (*usermodel.UserFind, error)
}

type findUserBiz struct {
	store FindUserStore
}

func NewFindUserBiz(store FindUserStore) *findUserBiz {
	return &findUserBiz{store: store}
}

func (b *findUserBiz) FindUser(ctx context.Context, userId int) (*usermodel.UserFind, error) {
	user, err := b.store.FindById(ctx, userId)

	if err != nil {
		if err == common.ErrorNoRows {
			return nil, common.ErrEntityNotFound(usermodel.EntityName, err)
		}
		return nil, err
	}

	if user.Status == common.UserDeleted {
		return nil, common.ErrEntityDeleted(usermodel.EntityName, err)
	}

	return user, nil
}
