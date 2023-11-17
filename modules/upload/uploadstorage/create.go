package uploadstorage

import (
	"bestHabit/common"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *common.Image) error {
	db := s.db
	query := `INSERT INTO images(url, width, height, cloud_name, extension) 
	VALUES(:url, :width, :height, :cloud_name, :extension)`

	if _, err := db.NamedExec(query, data); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
