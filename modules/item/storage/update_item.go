package storage

import (
	"context"
	"goobike-backend/modules/item/model"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	// Where id = (auto mapping)
	if err := s.db.Where(cond).Updates(&dataUpdate).Error; err != nil {
		if err != nil {
			return err
		}
	}

	return nil
}
