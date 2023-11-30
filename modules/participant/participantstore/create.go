package participantstore

import (
	"bestHabit/common"
	"bestHabit/modules/participant/participantmodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *participantmodel.ParticipantCreate) error {
	db := s.db

	query := `INSERT INTO participants (user_id, challenge_id) 
	VALUES (?,?)`

	if _, err := db.Exec(query, data.UserId, data.ChallengeId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) Rejoin(ctx context.Context, data *participantmodel.ParticipantCreate) error {
	db := s.db

	if _, err := db.Exec("UPDATE participants SET status = 'joined' WHERE user_id = ? AND challenge_id = ?",
		data.UserId, data.ChallengeId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
