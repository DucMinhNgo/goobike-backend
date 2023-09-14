package storage

import (
	"context"
	"goobike-backend/common"
	"goobike-backend/modules/item/model"

	"gorm.io/gorm"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	// Where id = (auto mapping)
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, common.ErrRecordNotFound
			}

			return nil, common.ErrDB(err)
		}
	}

	return &data, nil
}
