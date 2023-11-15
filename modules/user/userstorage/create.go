package userstorage

import (
	"bestHabit/common"
	"bestHabit/modules/user/usermodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db

	query := `
	INSERT INTO users (email, phone, password, name, fb_id, gg_id, salt, avatar, settings, role)
	VALUES (:email, :phone, :password, :name, :fb_id, :gg_id, :salt, :avatar, :settings, :role)`

	if _, err := db.NamedExec(query, data); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
