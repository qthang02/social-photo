package storage

import (
	"context"
	"social-photo/common"
	"social-photo/modules/userlikepost/model"
)

func (s *sqlStore) Delete(ctx context.Context, userId, postID int) error {
	var data model.Like

	if err := s.db.Table(data.TableName()).
		Where("user_id = ? and post_id = ?", userId, postID).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
