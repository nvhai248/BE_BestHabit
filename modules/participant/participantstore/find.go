package participantstore

import (
	"bestHabit/common"
	"bestHabit/modules/participant/participantmodel"
	"context"
	"database/sql"
)

func (s *sqlStore) FindParticipantById(ctx context.Context, id int) (*participantmodel.ParticipantFind, error) {
	db := s.db

	var participant participantmodel.ParticipantFind
	if err := db.Get(&participant, "SELECT * FROM participants WHERE id = ?", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrorNoRows
		}

		return nil, common.ErrDB(err)
	}

	return &participant, nil
}
