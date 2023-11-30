package participantbiz

import (
	"bestHabit/common"
	"bestHabit/modules/participant/participantmodel"
	"context"
)

type ListParticipantStore interface {
	ListParticipantByConditions(ctx context.Context,
		paging *common.Paging,
	) ([]participantmodel.Participant, error)
}

type listParticipantBiz struct {
	store ListParticipantStore
}

func NewListParticipantBiz(store ListParticipantStore) *listParticipantBiz {
	return &listParticipantBiz{store: store}
}

func (b *listParticipantBiz) ListChallengeJoined(ctx context.Context,
	paging *common.Paging) ([]participantmodel.Participant, error) {

	participants, err := b.store.ListParticipantByConditions(ctx, paging)

	if err != nil {
		return nil, err
	}

	return participants, nil
}
