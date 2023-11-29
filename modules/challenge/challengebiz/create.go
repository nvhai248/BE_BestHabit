package challengebiz

import (
	"bestHabit/modules/challenge/challengemodel"
	"context"
)

type CreateChallengeStore interface {
	Create(ctx context.Context, data *challengemodel.ChallengeCreate) error
}

type createChallengeBiz struct {
	store CreateChallengeStore
}

func NewCreateChallengeBiz(store CreateChallengeStore) *createChallengeBiz {
	return &createChallengeBiz{store: store}
}

func (b *createChallengeBiz) CreateChallenge(ctx context.Context, data *challengemodel.ChallengeCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := b.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
