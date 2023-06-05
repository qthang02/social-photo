package storage

import (
	"context"
	"social-photo/common"
	"social-photo/modules/post/model"
)

func (s *sqlStore) ListPost(ctx context.Context, paging *common.Paging, moreKey ...string) ([]model.Post, error) {

	var data []model.Post

	db := s.db

	if err := db.Table(model.Post{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Order("id desc").Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(data) > 0 {
		data[len(data)-1].Mask()
		paging.NextCursor = data[len(data)-1].FakeId.String()
	}

	return data, nil
}
