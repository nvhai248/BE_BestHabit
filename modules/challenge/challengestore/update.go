package challengestore

import (
	"bestHabit/common"
	"bestHabit/modules/challenge/challengemodel"
	"context"
)

func (s *sqlStore) UpdateChallengesInfo(ctx context.Context, newInfo *challengemodel.ChallengeUpdate, id int) error {
	db := s.db

	if _, err := db.Exec("UPDATE challenges SET name = ?, description = ?, start_date = ?, end_date = ?, experience_point = ? WHERE id = ?",
		newInfo.Name, newInfo.Description, newInfo.StartDate, newInfo.EndDate, newInfo.ExperiencePoint, id); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
