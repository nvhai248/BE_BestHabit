package participantbiz

import (
	"bestHabit/common"
	"bestHabit/modules/participant/participantmodel"
	"bestHabit/pubsub"
	"context"
)

type CancelParticipantStore interface {
	Cancel(ctx context.Context, data *participantmodel.ParticipantCancel) error
}

type cancelParticipantBiz struct {
	store CancelParticipantStore
	pb    pubsub.Pubsub
}

func NewCancelParticipantBiz(store CancelParticipantStore, pb pubsub.Pubsub) *cancelParticipantBiz {
	return &cancelParticipantBiz{store: store, pb: pb}
}

func (b *cancelParticipantBiz) CancelChallenge(ctx context.Context, data *participantmodel.ParticipantCancel) error {
	if err := b.store.Cancel(ctx, data); err != nil {
		return err
	}
	go func() {
		defer common.AppRecover()
		b.pb.Publish(ctx, common.TopicUserCancelChallenge, pubsub.NewMessage(data))
	}()
	return nil
}
