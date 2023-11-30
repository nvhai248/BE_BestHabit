package participantstore

import (
	"bestHabit/common"
	"bestHabit/modules/participant/participantmodel"
	"context"
)

func (s *sqlStore) UpdateParticipantInfo(ctx context.Context, newInfo *participantmodel.ParticipantUpdate, id int) error {
	db := s.db

	if _, err := db.Exec("UPDATE challenges SET status = ? WHERE id = ?",
		newInfo.Status, id); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
