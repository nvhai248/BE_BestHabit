package subscriber

import (
	"bestHabit/component"
	"bestHabit/modules/user/userstorage"
	"bestHabit/pubsub"
	"context"
)

func RunDecreaseTaskCountAfterUserDeleteTask(appCtx component.AppContext) consumerJob {
	store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Decrease task count after delete task!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			userData := message.Data().(HasUserId)
			return store.DecreaseTaskCount(ctx, userData.GetUserId())
		},
	}
}
