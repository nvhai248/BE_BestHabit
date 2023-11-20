package subscriber

import (
	"bestHabit/component"
	"bestHabit/modules/user/userstorage"
	"bestHabit/pubsub"
	"context"
)

func RunIncreaseTaskCountAfterUserCreateNewTask(appCtx component.AppContext) consumerJob {
	store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Increase task count after user create new task!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			userData := message.Data().(HasUserId)
			return store.IncreaseTaskCount(ctx, userData.GetUserId())
		},
	}
}
