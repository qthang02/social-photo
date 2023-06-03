package storage

import (
	"context"
	"social-photo/common"
	"social-photo/modules/userlikepost/model"
)

func (s *sqlStore) Create(ctx context.Context, data *model.Like) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
