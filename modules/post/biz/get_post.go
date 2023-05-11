package biz

import (
	"context"
	"social-photo/modules/post/model"
)

type GetPostStorage interface {
	GetPost(ctx context.Context, cond map[string]interface{}) (*model.Post, error)
}

type getPostBiz struct {
	store GetPostStorage
}

func NewGetPostBiz(store GetPostStorage) *getPostBiz {
	return &getPostBiz{store: store}
}

func (biz *getPostBiz) GetPostById(ctx context.Context, id int) (*model.Post, error) {
	data, err := biz.store.GetPost(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return data, nil
}
