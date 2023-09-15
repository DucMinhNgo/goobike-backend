package storage

import (
	"context"
	"goobike-backend/common"
	"goobike-backend/modules/user/model"

	"gorm.io/gorm"
)

func (s *sqlStore) GetUser(ctx context.Context, cond map[string]interface{}) (*model.User, error) {
	var data model.User

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
