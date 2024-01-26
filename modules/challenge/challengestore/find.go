package challengestore

import (
	"bestHabit/common"
	"bestHabit/modules/challenge/challengemodel"
	"context"
	"database/sql"
)

func (s *sqlStore) FindChallengesById(ctx context.Context, id int) (*challengemodel.ChallengeFind, error) {
	db := s.db

	var challenge challengemodel.ChallengeFind
	if err := db.Get(&challenge, "SELECT * FROM challenges WHERE id = ?", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrorNoRows
		}

		return nil, common.ErrDB(err)
	}

	return &challenge, nil
}

func (s *sqlStore) CountChallengesByTimeCreated(time string) (int, error) {
	db := s.db

	query := "select COUNT(id) from challenges where created_at LIKE '" + time + "%'"
	var count int
	if err := db.QueryRow(query).Scan(&count); err != nil {
		return 0, common.ErrDB(err)
	}

	return count, nil
}
