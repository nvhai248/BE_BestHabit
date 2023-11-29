package subscriber

import (
	"bestHabit/component"
	"bestHabit/modules/user/userstorage"
	"bestHabit/pubsub"
	"context"
)

func RunDecreaseHabitCountAfterUserDeleteHabit(appCtx component.AppContext) consumerJob {
	store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Decrease habit count after user delete habit!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			userData := message.Data().(HasUserId)
			return store.DecreaseHabitCount(ctx, userData.GetUserId())
		},
	}
}
