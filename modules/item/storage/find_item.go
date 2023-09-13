package storage

import (
	"context"
	"goobike-backend/modules/item/model"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	// Where id = (auto mapping)
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err != nil {
			return nil, err
		}
	}

	return &data, nil
}
