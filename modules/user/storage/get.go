package storage

import (
	"context"
	"gorm.io/gorm"
	"social-photo/common"
	"social-photo/modules/user/model"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}) (*model.User, error) {
	db := s.db.Table(model.User{}.TableName())

	var user model.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
