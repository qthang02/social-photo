package storage

import (
	"context"
	"social-photo/modules/post/model"
)

func (s *sqlStore) UpdatePost(ctx context.Context, cond map[string]interface{}, data *model.PostUpdate) error {
	if err := s.db.Where(cond).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
