package userstorage

import (
	"bestHabit/common"
	"bestHabit/modules/user/usermodel"
	"context"
	"database/sql"
)

func (s *sqlStore) FindById(ctx context.Context, id int) (*usermodel.UserFind, error) {
	db := s.db

	var result usermodel.UserFind

	if err := db.Get(&result, "SELECT * FROM users WHERE id = ?", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrorNoRows
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}

func (s *sqlStore) FindByEmail(ctx context.Context, email string) (*usermodel.UserFind, error) {
	db := s.db

	var result usermodel.UserFind

	if err := db.Get(&result, "SELECT * FROM users WHERE email = ?", email); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrorNoRows
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}

func (s *sqlStore) FindByGgId(ctx context.Context, ggId string) (*usermodel.UserFind, error) {
	db := s.db

	var result usermodel.UserFind

	if err := db.Get(&result, "SELECT * FROM users WHERE gg_id = ?", ggId); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrorNoRows
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}

func (s *sqlStore) FindByFbId(ctx context.Context, fbId string) (*usermodel.UserFind, error) {
	db := s.db

	var result usermodel.UserFind

	if err := db.Get(&result, "SELECT * FROM users WHERE fb_id = ?", fbId); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrorNoRows
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}

func (s *sqlStore) CountUserByTimeCreated(time string) (int, error) {
	db := s.db

	query := "select COUNT(id) from users where created_at LIKE '" + time + "%'"

	var count int
	if err := db.QueryRow(query).Scan(&count); err != nil {
		return 0, common.ErrDB(err)
	}

	return count, nil
}
