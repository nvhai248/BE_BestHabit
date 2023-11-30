package subscriber

import (
	"bestHabit/component"
	"bestHabit/modules/user/userstorage"
	"bestHabit/pubsub"
	"context"
)

func RunDecreaseChallengeCountAfterUserCancelChallenge(appCtx component.AppContext) consumerJob {
	store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Decrease challenge count after user cancel challenge!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			userData := message.Data().(HasUserId)
			return store.DecreaseChallengeCount(ctx, userData.GetUserId())
		},
	}
}
