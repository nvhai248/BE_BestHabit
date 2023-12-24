package taskbiz

import (
	"bestHabit/common"
	"bestHabit/component/cronjob"
	"bestHabit/modules/task/taskmodel"
	"bestHabit/pubsub"
	"context"
	"fmt"
)

type CreateTaskStore interface {
	Create(ctx context.Context, data *taskmodel.TaskCreate) error
}

type createTaskBiz struct {
	store  CreateTaskStore
	pubsub pubsub.Pubsub
}

func NewCreateTaskBiz(store CreateTaskStore, pubsub pubsub.Pubsub) *createTaskBiz {
	return &createTaskBiz{store: store, pubsub: pubsub}
}

func (b *createTaskBiz) CreateTask(ctx context.Context, data *taskmodel.TaskCreate, userId int) error {
	if err := data.Validate(); err != nil {
		return err
	}

	data.UserId = userId

	if err := b.store.Create(ctx, data); err != nil {
		return err
	}
	go func() {
		defer common.AppRecover()
		b.pubsub.Publish(ctx, common.TopicUserCreateNewTask, pubsub.NewMessage(data))
	}()

	// create a new cron job
	go func() {
		defer common.AppRecover()
		entryIds, _ := cronjob.CreateCronJob(*common.NewNotificationBasedOnTask(userId, data.Description, data.Name, data.Reminder))
		fmt.Println(entryIds)
	}()
	return nil
}
