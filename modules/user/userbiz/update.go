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

type updateProfileBiz struct {
	store UpdateProfileStore
}

func NewUpdateProfileBiz(store UpdateProfileStore) *updateProfileBiz {
	return &updateProfileBiz{store: store}
}

func (b *updateProfileBiz) UpdateProfile(ctx context.Context,
	newInfo *usermodel.UserUpdate,
	userId int) error {
	oldData, err := b.store.FindById(ctx, userId)
	if err != nil {
		if err == common.ErrorNoRows {
			return common.ErrEntityNotFound(usermodel.EntityName, err)
		}
		return common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.NewCustomError(nil, "Data deleted!", usermodel.EntityName)
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

	err = b.store.UpdateInfoById(ctx, newInfo, userId)

	if err != nil {
		return err
	}

	return nil
}
