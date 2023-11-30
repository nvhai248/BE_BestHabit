package participantbiz

import (
	"bestHabit/common"
	"bestHabit/modules/participant/participantmodel"
	"bestHabit/pubsub"
	"context"
)

type CreateParticipantStore interface {
	Create(ctx context.Context, data *participantmodel.ParticipantCreate) error
	FindParticipantByUserIdAndChallengeId(ctx context.Context, userId, challengeId int) (*participantmodel.ParticipantFind, error)
	Rejoin(ctx context.Context, data *participantmodel.ParticipantCreate) error
}

type createParticipantBiz struct {
	store CreateParticipantStore
	pb    pubsub.Pubsub
}

func NewCreateParticipantBiz(store CreateParticipantStore, pb pubsub.Pubsub) *createParticipantBiz {
	return &createParticipantBiz{store: store, pb: pb}
}

func (b *createParticipantBiz) CreateParticipant(ctx context.Context, data *participantmodel.ParticipantCreate) error {

	participant, err := b.store.FindParticipantByUserIdAndChallengeId(ctx, data.UserId, data.ChallengeId)

	if err != nil {
		if err != common.ErrorNoRows {
			return err
		} else {
			if err := b.store.Create(ctx, data); err != nil {
				return err
			}
		}
	} else {
		if participant.Status == "cancel" {
			err = b.store.Rejoin(ctx, data)

			if err != nil {
				return err
			}
		} else {
			return common.ErrEntityExisted(participantmodel.EntityName, err)
		}
	}

	go func() {
		defer common.AppRecover()
		b.pb.Publish(ctx, common.TopicUserJoinChallenge, pubsub.NewMessage(data))
	}()

	return nil
}
