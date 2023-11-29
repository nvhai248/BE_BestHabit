package challengebiz

import (
	"bestHabit/common"
	"bestHabit/modules/challenge/challengemodel"
	"context"
)

type ListChallengeStorage interface {
	ListChallengesByConditions(ctx context.Context,
		filter *challengemodel.ChallengeFilter,
		paging *common.Paging,
		conditions map[string]interface{}) ([]challengemodel.Challenge, error)
}

type listChallengeBiz struct {
	store ListChallengeStorage
}

func NewListChallengeBiz(store ListChallengeStorage) *listChallengeBiz {
	return &listChallengeBiz{store: store}
}

func (b *listChallengeBiz) ListChallenge(ctx context.Context,
	filter *challengemodel.ChallengeFilter,
	paging *common.Paging,
	conditions map[string]interface{}) ([]challengemodel.Challenge, error) {

	challenges, err := b.store.ListChallengesByConditions(ctx, filter, paging, conditions)

	if err != nil {
		return nil, err
	}

	return challenges, nil
}
