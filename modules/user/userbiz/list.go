package userbiz

import (
	"bestHabit/common"
	"bestHabit/modules/user/usermodel"
	"context"
)

type ListUserStorage interface {
	ListTaskByConditions(ctx context.Context,
		filter *usermodel.UserFilter,
		paging *common.Paging,
		conditions map[string]interface{}) ([]usermodel.User, error)
}

type listUserBiz struct {
	store ListUserStorage
}

func NewListUserBiz(store ListUserStorage) *listUserBiz {
	return &listUserBiz{store: store}
}

func (b *listUserBiz) ListTask(ctx context.Context,
	filter *usermodel.UserFilter,
	paging *common.Paging,
	conditions map[string]interface{}) ([]usermodel.User, error) {

	task, err := b.store.ListTaskByConditions(ctx, filter, paging, conditions)

	if err != nil {
		return nil, err
	}

	return task, nil
}
