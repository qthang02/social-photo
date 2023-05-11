package storage

import (
	"context"
	"social-photo/modules/post/model"
)

func (s *sqlStore) GetPost(ctx context.Context, cond map[string]interface{}) (*model.Post, error) {
	var data model.Post

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
