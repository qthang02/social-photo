package storage

import (
	"context"
	"gorm.io/gorm"
	"social-photo/common"
	"social-photo/modules/userlikepost/model"
)

func (s *sqlStore) Find(ctx context.Context, userId, postId int) (*model.Like, error) {
	var data model.Like

	if err := s.db.Where("user_id = ? and post_id = ?", userId, postId).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
