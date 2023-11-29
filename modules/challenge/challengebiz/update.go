package challengebiz

import (
	"bestHabit/common"
	"bestHabit/modules/challenge/challengemodel"
	"context"
)

type UpdateChallengeStorage interface {
	FindChallengesById(ctx context.Context, id int) (*challengemodel.ChallengeFind, error)
	UpdateChallengesInfo(ctx context.Context, newInfo *challengemodel.ChallengeUpdate, id int) error
}

type updateChallengeBiz struct {
	store UpdateChallengeStorage
}

func NewUpdateChallengeBiz(store UpdateChallengeStorage) *updateChallengeBiz {
	return &updateChallengeBiz{store: store}
}

func (b *updateChallengeBiz) Update(ctx context.Context, newInfo *challengemodel.ChallengeUpdate, id int) error {
	oldData, err := b.store.FindChallengesById(ctx, id)

	if err != nil {
		if err == common.ErrorNoRows {
			return common.ErrEntityNotFound(challengemodel.EntityName, err)
		}

		return err
	}

	if oldData.Status == false {
		return common.ErrEntityDeleted(challengemodel.EntityName, err)
	}

	if newInfo.Name == nil {
		newInfo.Name = &oldData.Name
	}

	if newInfo.Description == nil {
		newInfo.Description = &oldData.Description
	}

	if newInfo.StartDate == nil {
		newInfo.StartDate = &oldData.StartDate
	}

	if newInfo.EndDate == nil {
		newInfo.StartDate = &oldData.EndDate
	}

	if newInfo.ExperiencePoint == nil {
		newInfo.ExperiencePoint = &oldData.ExperiencePoint
	}

	err = b.store.UpdateChallengesInfo(ctx, newInfo, id)

	if err != nil {
		return common.ErrCannotUpdateEntity(challengemodel.EntityName, err)
	}

	return nil
}
