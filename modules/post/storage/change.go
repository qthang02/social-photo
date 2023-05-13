package storage

import (
	"context"
	"gorm.io/gorm"
	"social-photo/common"
	"social-photo/modules/post/model"
)

func (s *sqlStore) UpdatePost(ctx context.Context, cond map[string]interface{}, data *model.PostUpdate) error {
	if err := s.db.Where(cond).Updates(data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNotFound
		}

		return common.ErrDB(err)
	}

	return nil
}
