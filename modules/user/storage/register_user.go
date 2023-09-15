package storage

import (
	"context"
	"goobike-backend/common"
	"goobike-backend/modules/user/model"
)

func (s *sqlStore) RegisterUser(ctx context.Context, data *model.UserCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		// internal server error
		return common.ErrDB(err)
	}

	return nil
}
