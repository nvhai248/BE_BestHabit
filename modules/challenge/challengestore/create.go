package challengestore

import (
	"bestHabit/common"
	"bestHabit/modules/challenge/challengemodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *challengemodel.ChallengeCreate) error {
	db := s.db

	query := `INSERT INTO challenges (name, description, start_date, end_date, experience_point) 
	VALUES (?,?,?,?,?)`

	if _, err := db.Exec(query, data.Name, data.Description, data.StartDate, data.EndDate, data.ExperiencePoint); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
