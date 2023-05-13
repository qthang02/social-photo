package storage

import (
	"context"
	"social-photo/common"
	"social-photo/modules/post/model"
)

func (s *sqlStore) ListPost(ctx context.Context, paging *common.Paging) ([]model.Post, error) {

	var data []model.Post

	if err := s.db.Table(model.Post{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := s.db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
