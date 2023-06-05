package storage

import (
	"context"
	"social-photo/common"
	"social-photo/modules/userlikepost/model"
)

func (s *sqlStore) ListUser(ctx context.Context, postId int, paging *common.Paging) ([]common.SimpleUser, error) {
	var result []model.Like

	db := s.db.Where("post_id = ?", postId)

	if err := db.Table(model.Like{}.TableName()).Select("user_id").Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Select("*").
		Order("created_at desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Preload("User").
		Find(&result).Error; err != nil {

		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))

	for i := range users {
		users[i] = *result[i].User
		users[i].UpdatedAt = nil
		users[i].CreatedAt = result[i].CreatedAt
	}

	return users, nil
}
