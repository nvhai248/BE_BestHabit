package subscriber

import (
	"bestHabit/component"
	"bestHabit/modules/challenge/challengestore"
	"bestHabit/pubsub"
	"context"
)

func RunDecreaseUserJoinedCountAfterUserCancelChallenge(appCtx component.AppContext) consumerJob {
	store := challengestore.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Decrease user joined count after user cancel challenge!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			userData := message.Data().(HasChallengeId)
			return store.DecreaseCountUserJoined(ctx, userData.GetChallengeId())
		},
	}
}
