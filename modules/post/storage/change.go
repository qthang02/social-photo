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

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(model.Post{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(model.Post{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
