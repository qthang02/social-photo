package storage

import (
	"context"
	"gorm.io/gorm"
	"social-photo/common"
	"social-photo/modules/post/model"
)

func (s *sqlStore) GetPost(ctx context.Context, cond map[string]interface{}) (*model.Post, error) {
	var data model.Post

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
