package participantbiz

import (
	"bestHabit/common"
	"bestHabit/modules/participant/participantmodel"
	"context"
)

type FindParticipantStore interface {
	FindParticipantById(ctx context.Context, id int) (*participantmodel.ParticipantFind, error)
}

type findParticipantBiz struct {
	store FindParticipantStore
}

func NewFindParticipantBiz(store FindParticipantStore) *findParticipantBiz {
	return &findParticipantBiz{store: store}
}

func (b *findParticipantBiz) FindChallengeJoined(ctx context.Context, id int) (*participantmodel.ParticipantFind, error) {
	result, err := b.store.FindParticipantById(ctx, id)

	if err != nil {
		if err == common.ErrorNoRows {
			return nil, common.ErrCannotGetEntity(participantmodel.EntityName, err)
		}
		return nil, err
	}

	if result.Status == "cancel" {
		return nil, common.ErrEntityDeleted(participantmodel.EntityName, err)
	}

	return result, nil
}
