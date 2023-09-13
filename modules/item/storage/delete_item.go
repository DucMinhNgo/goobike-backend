package storage

import (
	"context"
	"goobike-backend/modules/item/model"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := model.ItemStatusDeleted
	// hard delete
	// if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
	// soft delete
	if err := s.db.Table(model.TodoItem{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"status": deletedStatus.String(),
	}).Error; err != nil {
		if err != nil {
			return err
		}
	}

	return nil
}
