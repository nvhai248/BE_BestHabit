package subscriber

import (
	"bestHabit/component"
	"bestHabit/modules/challenge/challengestore"
	"bestHabit/pubsub"
	"context"
)

func RunIncreaseUserJoinedCountAfterUserJoinedChallenge(appCtx component.AppContext) consumerJob {
	store := challengestore.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Increase user joined count after user cancel challenge!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			userData := message.Data().(HasChallengeId)
			return store.IncreaseCountUserJoined(ctx, userData.GetChallengeId())
		},
	}
}
