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

func (s *sqlStore) FindParticipantByUserIdAndChallengeId(ctx context.Context, userId, challengeId int) (*participantmodel.ParticipantFind, error) {
	db := s.db

	var participant participantmodel.ParticipantFind
	if err := db.Get(&participant, "SELECT * FROM participants WHERE user_id = ? AND challenge_id = ?", userId, challengeId); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrorNoRows
		}

		return nil, common.ErrDB(err)
	}

	return &participant, nil
}
