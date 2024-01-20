package userstorage

import (
	"bestHabit/common"
	"bestHabit/modules/user/usermodel"
	"context"
)

func (s *sqlStore) IncreaseHabitCount(ctx context.Context, userId int) error {
	db := s.db

	if _, err := db.Exec("UPDATE users SET habit_count = habit_count + 1 WHERE id = ?", userId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseTaskCount(ctx context.Context, userId int) error {
	db := s.db

	if _, err := db.Exec("UPDATE users SET task_count = task_count + 1 WHERE id = ?", userId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseChallengeCount(ctx context.Context, userId int) error {
	db := s.db

	if _, err := db.Exec("UPDATE users SET challenge_count = challenge_count + 1 WHERE id = ?", userId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) UpdateInfoById(ctx context.Context,
	newInfo *usermodel.UserUpdate,
	userId int) error {
	db := s.db

	if _, err := db.Exec("UPDATE users SET name = ?, phone = ?, avatar = ?, settings = ? WHERE id = ?",
		newInfo.Name, newInfo.Phone, newInfo.Avatar, newInfo.Settings, userId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseHabitCount(ctx context.Context, userId int) error {
	db := s.db

	if _, err := db.Exec("UPDATE users SET habit_count = habit_count - 1 WHERE id = ?", userId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseTaskCount(ctx context.Context, userId int) error {
	db := s.db

	if _, err := db.Exec("UPDATE users SET task_count = task_count - 1 WHERE id = ?", userId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseChallengeCount(ctx context.Context, userId int) error {
	db := s.db

	if _, err := db.Exec("UPDATE users SET challenge_count = challenge_count - 1 WHERE id = ?", userId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) VerifyUser(ctx context.Context, userId int) error {
	db := s.db

	if _, err := db.Exec("UPDATE users SET status = 1 WHERE id = ?", userId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) ChangePassword(ctx context.Context, newPw string, userId int) error {
	db := s.db

	if _, err := db.Exec("UPDATE users SET password = ? WHERE id = ?", newPw, userId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) BannedUser(ctx context.Context, userId int) error {
	db := s.db

	if _, err := db.Exec("UPDATE users SET status = -1 WHERE id = ?", userId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) UnbannedUser(ctx context.Context, userId int) error {
	db := s.db

	if _, err := db.Exec("UPDATE users SET status = 1 WHERE id = ?", userId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
