package taskbiz

import (
	"bestHabit/common"
	"bestHabit/modules/task/taskmodel"
	"bestHabit/pubsub"
	"context"
)

type DeleteTaskStore interface {
	DeleteTask(ctx context.Context, id int) error
	FindTaskById(ctx context.Context, id int) (*taskmodel.TaskFind, error)
}

type deleteTaskBiz struct {
	store  DeleteTaskStore
	pubsub pubsub.Pubsub
}

func NewDeleteTaskBiz(store DeleteTaskStore, pubsub pubsub.Pubsub) *deleteTaskBiz {
	return &deleteTaskBiz{store: store, pubsub: pubsub}
}

func (b *deleteTaskBiz) SoftDeleteTask(ctx context.Context, id int) error {

	task, err := b.store.FindTaskById(ctx, id)

	if err != nil {
		if err == common.ErrorNoRows {
			return common.ErrEntityNotFound(taskmodel.EntityName, err)
		}

		return common.ErrCannotGetEntity(taskmodel.EntityName, err)
	}

	if task.Status == "deleted" {
		return common.ErrCannotDeleteEntity(taskmodel.EntityName, err)
	}

	if err := b.store.DeleteTask(ctx, id); err != nil {
		return err
	}
	go func() {
		defer common.AppRecover()
		b.pubsub.Publish(ctx, common.TopicUserDeleteTask, pubsub.NewMessage(task))
	}()
	return nil
}
