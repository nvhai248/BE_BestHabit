package uploadstorage

import (
	"bestHabit/common"
	"bestHabit/modules/upload/uploadmodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *uploadmodel.ImageUpload) error {
	db := s.db
	query := `INSERT INTO images(url, width, height, cloud_name, extension, created_by) 
	VALUES(:url, :width, :height, :cloud_name, :extension, :created_by)`

	if _, err := db.NamedExec(query, data); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
