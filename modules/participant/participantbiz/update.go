package participantbiz

import (
	"bestHabit/common"
	"bestHabit/modules/participant/participantmodel"
	"context"
)

type UpdateParticipantStore interface {
	FindParticipantById(ctx context.Context, id int) (*participantmodel.ParticipantFind, error)
	UpdateParticipantInfo(ctx context.Context, newInfo *participantmodel.ParticipantUpdate, id int) error
}

type updateParticipantBiz struct {
	store UpdateParticipantStore
}

func NewUpdateParticipantBiz(store UpdateParticipantStore) *updateParticipantBiz {
	return &updateParticipantBiz{store: store}
}

func (b *updateParticipantBiz) Update(ctx context.Context, newInfo *participantmodel.ParticipantUpdate, id int) error {
	oldData, err := b.store.FindParticipantById(ctx, id)

	if err != nil {
		if err == common.ErrorNoRows {
			return common.ErrEntityNotFound(participantmodel.EntityName, err)
		}

		return err
	}

	if oldData.Status == "cancel" {
		return common.ErrEntityDeleted(participantmodel.EntityName, err)
	}

	err = b.store.UpdateParticipantInfo(ctx, newInfo, id)

	if err != nil {
		return common.ErrCannotUpdateEntity(participantmodel.EntityName, err)
	}

	return nil
}
