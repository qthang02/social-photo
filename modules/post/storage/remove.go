package storage

import (
	"context"
	"social-photo/modules/post/model"
)

func (s *sqlStore) DeletePost(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(model.Post{}.TableName()).Where(cond).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}