package storage

import (
	"context"
	"goobike-backend/common"
	"goobike-backend/modules/item/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		// internal server error
		return common.ErrDB(err)
	}

	return nil
}
