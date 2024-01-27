package userbiz

import (
	"bestHabit/common"
	"bestHabit/modules/user/usermodel"
	"context"
)

type UpdateDVTokenStore interface {
	UpdateDeviceToken(ctx context.Context, userId int, data *usermodel.UpdateDeviceTokens) error

	FindById(ctx context.Context, id int) (*usermodel.UserFind, error)
}

type updateDVTokenBiz struct {
	store UpdateDVTokenStore
}

func NewUpdateDVTokenBiz(store UpdateDVTokenStore) *updateDVTokenBiz {
	return &updateDVTokenBiz{store: store}
}

func (b *updateDVTokenBiz) UpdateDVToken(ctx context.Context,
	newInfo *common.DvToken,
	userId int) error {
	oldData, err := b.store.FindById(ctx, userId)
	if err != nil {
		if err == common.ErrorNoRows {
			return common.ErrEntityNotFound(usermodel.EntityName, err)
		}
		return common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	if oldData.Status == common.UserDeleted {
		return common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	if oldData.DeviceTokens == nil {
		oldData.DeviceTokens = &common.DvTokens{}
	}

	oldData.DeviceTokens.AddNewDvToken(*newInfo)

	err = b.store.UpdateDeviceToken(ctx, userId, &usermodel.UpdateDeviceTokens{
		DeviceTokens: oldData.DeviceTokens,
	})

	if err != nil {
		return err
	}

	return nil
}
