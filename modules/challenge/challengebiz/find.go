package challengebiz

import (
	"bestHabit/common"
	"bestHabit/modules/challenge/challengemodel"
	"bestHabit/modules/task/taskmodel"
	"context"
)

type FindChallengeStorage interface {
	FindChallengesById(ctx context.Context, id int) (*challengemodel.ChallengeFind, error)
}

type findChallengeBiz struct {
	store FindChallengeStorage
}

func NewFindChallengeBiz(store FindChallengeStorage) *findChallengeBiz {
	return &findChallengeBiz{store: store}
}

func (b *findChallengeBiz) FindChallenge(ctx context.Context, id int) (*challengemodel.ChallengeFind, error) {
	result, err := b.store.FindChallengesById(ctx, id)

	if err != nil {
		if err == common.ErrorNoRows {
			return nil, common.ErrCannotGetEntity(taskmodel.EntityName, err)
		}
		return nil, err
	}

	if result.Status == false {
		return nil, common.ErrEntityDeleted(taskmodel.EntityName, err)
	}

	return result, nil
}
