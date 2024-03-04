package userbiz

import (
	"bestHabit/common"
	"bestHabit/modules/user/usermodel"
	"context"
)

type UpdateProfileStore interface {
	UpdateInfoById(ctx context.Context,
		newInfo *usermodel.UserUpdate,
		userId int) error

	FindById(ctx context.Context, id int) (*usermodel.UserFind, error)
}

type UpdateUserInfo interface {
	UpdateUserInfoByGRPC(ctx context.Context, userId int, userUpdate *usermodel.UserUpdate) (*usermodel.User, error)
}

type updateProfileBiz struct {
	store          UpdateProfileStore
	updateUserInfo UpdateUserInfo
}

func NewUpdateProfileBiz(store UpdateProfileStore, updateUserInfo UpdateUserInfo) *updateProfileBiz {
	return &updateProfileBiz{store: store, updateUserInfo: updateUserInfo}
}

func (b *updateProfileBiz) UpdateProfile(ctx context.Context,
	newInfo *usermodel.UserUpdate,
	userId int) (*usermodel.User, error) {
	oldData, err := b.store.FindById(ctx, userId)
	if err != nil {
		if err == common.ErrorNoRows {
			return nil, common.ErrEntityNotFound(usermodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	if oldData.Status == common.UserDeleted {
		return nil, common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	if newInfo.Name == nil {
		newInfo.Name = oldData.Name
	}

	if newInfo.Phone == nil {
		newInfo.Phone = oldData.Phone
	}

	if newInfo.Avatar == nil {
		newInfo.Avatar = oldData.Avatar
	}

	if newInfo.Settings == nil {
		newInfo.Settings = oldData.Settings
	}

	/* err = b.store.UpdateInfoById(ctx, newInfo, userId) */

	newData, err := b.updateUserInfo.UpdateUserInfoByGRPC(ctx, userId, newInfo)

	if err != nil {
		return nil, err
	}

	return newData, nil
}
