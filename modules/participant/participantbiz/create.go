package participantbiz

import (
	"bestHabit/common"
	"bestHabit/modules/participant/participantmodel"
	"bestHabit/pubsub"
	"context"
)

type CreateParticipantStore interface {
	Create(ctx context.Context, data *participantmodel.ParticipantCreate) error
}

type createParticipantBiz struct {
	store CreateParticipantStore
	pb    pubsub.Pubsub
}

func NewCreateParticipantBiz(store CreateParticipantStore, pb pubsub.Pubsub) *createParticipantBiz {
	return &createParticipantBiz{store: store, pb: pb}
}

func (b *createParticipantBiz) CreateParticipant(ctx context.Context, data *participantmodel.ParticipantCreate) error {

	if err := b.store.Create(ctx, data); err != nil {
		return err
	}
	go func() {
		defer common.AppRecover()
		b.pb.Publish(ctx, common.TopicUserJoinChallenge, pubsub.NewMessage(data))
	}()
	return nil
}
