package taskbiz

import (
	"bestHabit/common"
	"bestHabit/modules/task/taskmodel"
	"bestHabit/pubsub"
	"context"
	"fmt"
)

type UpdateTaskStorage interface {
	FindTaskById(ctx context.Context, id int) (*taskmodel.TaskFind, error)
	UpdateTaskInfo(ctx context.Context, newInfo *taskmodel.TaskUpdate, id int) error
}

type updateTaskBiz struct {
	store  UpdateTaskStorage
	pubsub pubsub.Pubsub
}

func NewUpdateTaskBiz(store UpdateTaskStorage, pubsub pubsub.Pubsub) *updateTaskBiz {
	return &updateTaskBiz{store: store, pubsub: pubsub}
}

func (b *updateTaskBiz) Update(ctx context.Context, newInfo *taskmodel.TaskUpdate, id int) error {
	oldData, err := b.store.FindTaskById(ctx, id)

	isNeedUpdateCronJob := true

	if err != nil {
		if err == common.ErrorNoRows {
			return common.ErrEntityNotFound(taskmodel.EntityName, err)
		}

		return err
	}

	if oldData.Status == "deleted" {
		return common.ErrEntityDeleted(taskmodel.EntityName, err)
	}

	if newInfo.Status == nil {
		newInfo.Status = &oldData.Status
	}

	if newInfo.Name == nil {
		newInfo.Name = &oldData.Name
	}

	if newInfo.Description == nil {
		newInfo.Description = &oldData.Description
	}

	if newInfo.Deadline == nil {
		newInfo.Deadline = &oldData.Deadline
	}

	if newInfo.Reminder == nil {
		newInfo.Reminder = &oldData.Reminder
		isNeedUpdateCronJob = false
	}

	newInfo.UserId = &oldData.UserId
	newInfo.Id = &id

	err = b.store.UpdateTaskInfo(ctx, newInfo, id)

	if err != nil {
		return common.ErrCannotUpdateEntity(taskmodel.EntityName, err)
	}

	fmt.Println(isNeedUpdateCronJob)

	if isNeedUpdateCronJob {
		go func() {
			defer common.AppRecover()
			b.pubsub.Publish(ctx, common.TopicUserUpdateTask, pubsub.NewMessage(newInfo))
		}()
	}

	return nil
}
