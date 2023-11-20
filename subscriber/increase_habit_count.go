package subscriber

import (
	"bestHabit/component"
	"bestHabit/modules/user/userstorage"
	"bestHabit/pubsub"
	"context"
)

func RunIncreaseHabitCountAfterUserCreateNewHabit(appCtx component.AppContext) consumerJob {
	store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Increase habit count after user create new habit!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			userData := message.Data().(HasUserId)
			return store.IncreaseHabitCount(ctx, userData.GetUserId())
		},
	}
}
