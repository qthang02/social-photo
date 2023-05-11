package storage

import (
	"context"
	"social-photo/modules/post/model"
)

func (s *sqlStore) CreatePost(ctx context.Context, data *model.PostCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
