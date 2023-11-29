package challengebiz

import (
	"context"
)

type DeleteChallengeStore interface {
	SoftDelete(ctx context.Context, id int) error
}

type deleteChallengeBiz struct {
	store DeleteChallengeStore
}

func NewDeleteChallengeBiz(store DeleteChallengeStore) *deleteChallengeBiz {
	return &deleteChallengeBiz{store: store}
}

func (b *deleteChallengeBiz) SoftDeleteChallenge(ctx context.Context, id int) error {
	if err := b.store.SoftDelete(ctx, id); err != nil {
		return err
	}

	return nil
}
