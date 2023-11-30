package participantstore

import (
	"bestHabit/common"
	"bestHabit/modules/participant/participantmodel"
	"context"
)

func (s *sqlStore) Cancel(ctx context.Context, data *participantmodel.ParticipantCancel) error {
	db := s.db

	if _, err := db.Exec("UPDATE participants SET status = 'cancel' WHERE user_id = ? AND challenge_id = ?",
		data.UserId, data.ChallengeId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
